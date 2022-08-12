package main

// Import packages
import (
	"fmt"
	"strings"
	"unicode"

	"github.com/valyala/fasthttp"
)

////////////////////////////////////////////////////////////
//														  //
//														  //
//														  //
// 				   Subject Code Scraping 	     		  //
//														  //
//														  //
//														  //
////////////////////////////////////////////////////////////

// The _ScrapSubjectCodes() function will return a slice containing all
// the course codes from the provided html
// The html is from: https://classes.uwaterloo.ca/uwpcshtm.html
//
// The function takes the html string pointer which is the websites
// string response body
//
// The function returns the string slice of scraped subject codes
func _ScrapeSubjectCodes(html *string) []string {
	// Declare Variables
	// res: []string -> result slice holding the subject codes
	// tableIndex: int -> used to only append the subject codes to res
	var (
		res        []string = []string{}
		tableIndex int      = 1
	)
	// Set the html to past the </table>
	html = &strings.Split(*html, "</table>")[1]

	// Iterate over the split strings
	for i, p := range strings.Split(*html, "<td>") {
		// Get every 7th table value (the Subject table)
		if i == tableIndex {
			// Increase tableIndex by 7
			tableIndex += 7
			// Split by closing tag
			var s []string = strings.Split(p, "</td>")
			// If the result slice doesn't contains the subject
			if !SliceContains(&res, s[0]) {
				// Append the subject to the result slice
				res = append(res, s[0])
			}
		}
	}
	// Return the result slice
	return res
}

// The ScrapSubjectCodes() function utilizes the _ScrapeSubjectCodes()
// function to webscrape all the subject codes on the
// "https://classes.uwaterloo.ca/uwpcshtm.html" website
//
// The Subject Codes will be used to get information about the courses
// from the "https://ucalendar.uwaterloo.ca/2021/COURSE/course-CS.html"
// website.
//
// The function takes the client: *fasthttp.Client parameter to send
// http requests
//
// The function returns the scraped subject codes string slice
// and the http request error
func ScrapeSubjectCodes(client *fasthttp.Client) ([]string, error) {
	// Utilize the _Request struct to easily send an http request
	var _Req *HttpRequest = &HttpRequest{
		Client: client,
		Url:    "https://classes.uwaterloo.ca/uwpcshtm.html",
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
		},
	}
	// Send Http Request
	var resp, err = _Req.Send()

	// Handle any request errors
	if err != nil {
		return []string{}, err
	}

	// Define response body variable
	var body string = string(resp.Body())

	// Scrape the subject codes using the response.body()
	// Then return the codes alongside no error
	return _ScrapeSubjectCodes(&body), nil
}

////////////////////////////////////////////////////////////
//														  //
//														  //
//														  //
// 				   Course Data Scraping 	     		  //
//														  //
//														  //
//														  //
////////////////////////////////////////////////////////////

// Course Scrape struct to help organize the data
// The Course Scrape struct holds three keys
// - Index: int -> used to select which Result key to use
// - Data: []string -> the index data
// - Result: map[string]string -> the course data result map
type CourseScrape struct {
	Index  int
	Data   []string
	Result map[string]string
}

// Convert the course course info into categories
// For example it will convert MATH 235 LEC,IST,TUT 0.50
// to -> var courseTitle, components, unit = MATH 235, LEC,IST,TUT, 0.50
//
// The function takes the title: string paramater
//
// The function returns the course title string,
// the components string and the unit string
func IndexCourseInfo(title string) (string, string, string) {
	// splitTitle: []string -> splits the given title by spaces
	// courseTitle: string -> the course name and catalog number (MATH 232)
	var (
		splitTitle  []string = strings.Split(title, " ")
		courseTitle string   = fmt.Sprintf("%s %s", splitTitle[0], splitTitle[1])
	)
	// Return the Course Title ex: CS 201
	// Return the Components (splitTitle[2]) ex: LAB,LEC,TST
	// Return the Unit (splitTitle[3]) ex: 0.50
	return courseTitle, splitTitle[2], splitTitle[3]
}

// The SetCourseInfo() function takes in the original course info
// and sets the title, components and unit using the IndexCourseInfo
// function. This makes it much easier to read on the front end
//
// The function takes the cs: *CourseScrape parameter
func SetCourseInfo(cs *CourseScrape) {
	if len(cs.Data) > 0 {
		cs.Result["title"],
			cs.Result["components"],
			cs.Result["unit"] = IndexCourseInfo(cs.Data[0])
	}
}

// The SetCourseId() function will set the CourseScrape struct object
// Result map id. The Course id is the key used to categorize all the data
// from each course. Example: id: {data: values}
//
// The function takes the cs: *CourseScrape parameter
func SetCourseId(cs *CourseScrape) {
	// Data Length is greater than one
	if len(cs.Data) > 1 {
		// Split the string by "Course ID: " to get the exact id
		var split []string = strings.Split(cs.Data[1], "Course ID: ")
		// Make sure the split length is greater than one
		if len(split) > 1 {
			// Set the result map value
			cs.Result["id"] = split[1]
		}
	}
}

// The SetCourseName() function will set the CourseScrape struct object
// Result map name.
//
// The function takes the cs: *CourseScrape parameter
func SetCourseName(cs *CourseScrape) {
	if len(cs.Data) > 2 {
		cs.Result["name"] = cs.Data[2]
	}
}

// The SetCourseDescription() function will set the CourseScrape struct object
// Result map description.
//
// The function takes the cs: *CourseScrape parameter
func SetCourseDescription(cs *CourseScrape) {
	if len(cs.Data) > 1 {
		cs.Result["desc"] = cs.Data[1]
	}
}

// The SetCourseAnti_CoReqs() function will set the CourseScrape struct object
// Result map anti_reqs key or the co_reqs key pr the pre_reqs
//
// The Anti Reqs key, Co Reqs key and the pre reqs key are in the same function because
// Sometimes the university of waterloo website will have Pre Requisites,
// or Co Requisites instead of Anti-Requisites, vice versa
//
// The function takes the cs: *CourseScrape parameter
func SetCourseAnti_Co_PreReqs(cs *CourseScrape) {
	if len(cs.Data) > 2 {
		// Start with anti reqs
		var splitBy, key string = "Antireq: ", "anti_reqs"

		// If it shows coreqs instead of antireqs, change the names
		if strings.Contains(cs.Data[2], "Coeq: ") {
			splitBy, key = "Coreq: ", "co_reqs"

			// If it shows prereqs instead of antireqs, change the names
		} else if strings.Contains(cs.Data[2], "Prereq: ") {
			splitBy, key = "Prereq: ", "pre_reqs"
		}

		// Split the string by the name_1
		var split []string = strings.Split(cs.Data[2], splitBy)
		if len(split) > 1 {
			cs.Result[key] = split[1]
		}
	}
}

// The IndexCourseScrapeResult() uses a map to categorize all the functions
// that will be used for indexing the scrap result
//
// I decided to go with this way because it was cleaner than having
// a bunch of if and else if statements
//
// I also decided to use an int index for categorizing everything instead
// of having to call if strings.Contains() {} a bunch of times
//
// The function takes the cs: *CourseScrape parameter
//
// The function returns the CourseScrape Result map pointer
func IndexCourseScrapeResult(cs *CourseScrape) *map[string]string {
	// The map for the CourseScrape.Index
	var indexMap map[int]func(cs *CourseScrape) = map[int]func(cs *CourseScrape){
		1: SetCourseInfo,            // title, components, unit
		2: SetCourseId,              // id
		3: SetCourseName,            // name
		4: SetCourseDescription,     // desc
		5: func(_ *CourseScrape) {}, // [empty]
		6: SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		7: SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		8: SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
	}
	// Call the function
	indexMap[cs.Index](cs)

	// Return result map
	return &cs.Result
}

// The _ScrapeCourseData() function will create a result map
// that stores the course data. The course data map holds the
// course id, name, description, pre-reqs, anti-reqs, etc.
//
// The function takes the table: *string parameter
//
// The function returns the course id string and the result map[string]string
func _ScrapeCourseData(table *string) (string, map[string]string) {
	// Define Variables
	var (
		// Create a CourseScrap struct object
		cs *CourseScrape = &CourseScrape{
			Result: make(map[string]string),
			Index:  0,
		}
		// Split the table into the segments that contain the course info
		splitTable []string = strings.Split(*table, "</")[1:]
	)

	// Iterate through the split table
	for i := 0; i < len(splitTable); i++ {
		// Split the splitted table by >
		cs.Data = strings.Split(splitTable[i], ">")[1:]

		// Check Data length
		if len(cs.Data[0]) > 1 {
			// Check if the splitTable contains a note about the course
			if !strings.Contains(splitTable[i], "Note:") {
				cs.Index++
				// Break the loop if the index is higher than 8
				if cs.Index > 8 {
					break
				}
				// Index the scrape result
				cs.Result = *IndexCourseScrapeResult(cs)
			} else {
				// Split the string to get note content
				var split []string = strings.Split(splitTable[i], "[Note: ")
				if len(split) > 1 {
					// Set the note in the result map
					cs.Result["note"] = split[1]
				}
			}
		}
	}
	// Return the course id and the course info map (result)
	return cs.Result["id"], cs.Result
}

// The CleaCourseTitle() function will remove all spaces and
// new lines from the h2 header
//
// The function returns the cleaned title string
func CleanCourseTitle(title string) string {
	var res string = ""
	for i := 0; i < len(title); i++ {
		// If the value == "&" increase by four to avoid
		// the &nmbp string
		if title[i] == '&' {
			i += 4
		} else if unicode.IsLetter(rune(title[i])) {
			// Append the letter to the result string
			res += string(title[i])
		}
	}
	// Return the res string in lowercase
	return strings.ToLower(res)
}

// The _ScrapeCourseTitle() function will return the title of the course at
// the top of the "https://classes.uwaterloo.ca/uwpcshtm.html" website
//
// This title will be used for search indexing thus it needs to be cleaned using
// the CleanCourseTitle() function
//
// The function uses the scraped title to return the cleaned course title string
func _ScrapeCourseTitle(body *string) string {
	// Define variables -> Getting the course title string
	var (
		_title string = strings.Split(*body, "<h2 class=\"subject\">")[1]
		title  string = strings.Split(_title, "</h2>")[0]
	)

	// Clean the course title and return it
	// Example: C O M P U T E R -> computer
	return CleanCourseTitle(title)
}

// The ScrapeCourseData() function is the main course scraper function
// This is because it scrapes all the course information and appends
// it to a map
//
// The function takes the client: *fasthttp.Client parameter to send http requests
//
// The function returns the course title string, the course data result map
// and the http request error
func ScrapeCourseData(client *fasthttp.Client, course string) (string, *map[string]map[string]string, error) {
	// Utilize the HttpRequest struct to easily send an http request
	var _Req *HttpRequest = &HttpRequest{
		Client: client,
		Url:    fmt.Sprintf("https://ucalendar.uwaterloo.ca/2021/COURSE/course-%s.html", course),
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
		},
	}
	// Define Variables
	// resp, err -> request response and error
	// result: map[string]map[string]string -> The result map that holds all the course data
	var (
		resp, err = _Req.Send()
		result    = make(map[string]map[string]string)
	)

	// Handle response error
	if err != nil {
		return "", &result, err
	}

	// Define Variables
	// body: string -> The http response body
	// courseTables: []string -> The tables with each course program data
	// courseTitle: string -> The courses title (ex: CS -> computerscience)
	var (
		body         string   = string(resp.Body())
		courseTables []string = strings.Split(body, "<div class=\"divTable\">")[1:]
		courseTitle  string   = _ScrapeCourseTitle(&body)
	)

	// Iterate over the html tables
	for _, table := range courseTables {
		var courseID, courseInfo = _ScrapeCourseData(&table)
		result[courseID] = courseInfo
	}

	// Return the course title and it's result map containing
	// all the courses information
	return courseTitle, &result, nil
}
