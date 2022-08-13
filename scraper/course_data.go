package scraper

// Import packages
import (
	"fmt"
	"strings"
	"sync"

	http "github.com/realTristan/The_University_of_Waterloo/http"
	"github.com/valyala/fasthttp"
)

// The Course Scrape struct holds three keys
/* - Index: int -> used to select which Result key to use 		*/
/* - Row: []string -> the row with data							*/
/* - Result: map[string]string -> the course data result map	*/
type CourseScrape struct {
	Index  int
	Row    []string
	Result map[string]string
}

// Convert the course course info into categories
// For example it will convert MATH 235 LEC,IST,TUT 0.50 to
// var courseTitle, components, unit = MATH 235, LEC,IST,TUT, 0.50
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
	if len(cs.Row) > 0 {
		cs.Result["title"],
			cs.Result["components"],
			cs.Result["unit"] = IndexCourseInfo(cs.Row[0])
	}
}

// The SetCourseId() function will set the CourseScrape
// Result map id. The Course id is the key used to categorize all the data
// from each course. Example: id: {data: values}
//
// The function takes the cs: *CourseScrape parameter
func SetCourseId(cs *CourseScrape) {
	// Data Length is greater than one
	if len(cs.Row) > 1 {
		// Split the string by "Course ID: " to get the exact id
		var split []string = strings.Split(cs.Row[1], "Course ID: ")
		// Make sure the split length is greater than one
		if len(split) > 1 {
			// Set the result map value
			cs.Result["id"] = split[1]
		}
	}
}

// The SetCourseName() function will set the CourseScrape
// Result map name -> Example result["name"] = "Intro to Computer Science"
//
// The function takes the cs: *CourseScrape parameter
func SetCourseName(cs *CourseScrape) {
	if len(cs.Row) > 2 {
		cs.Result["name"] = cs.Row[2]
	}
}

// The SetCourseDescription() function will set the CourseScrape
// Result map description -> Example result["desc"] = "Learn about..."
//
// The function takes the cs: *CourseScrape parameter
func SetCourseDescription(cs *CourseScrape) {
	if len(cs.Row) > 1 {
		cs.Result["desc"] = cs.Row[1]
	}
}

// The SetCourseNote() function will set the CourseScrape
// Result map note -> Example result["note"] = "Only available for..."
//
// The function takes the cs: *CourseScrape parameter
// and the data string parameter which is the html row of the note
func SetCourseNote(cs *CourseScrape, data string) {
	// Split the string to get note content
	var split []string = strings.Split(data, "[Note: ")
	if len(split) > 1 {
		// Set the note in the result map
		cs.Result["note"] = split[1]
	}
}

// The SetCourseAnti_Co_PreReqs() function will set the CourseScrape
// Result map anti_reqs, co_reqs or pre_reqs key
//
// The Anti Reqs key, Co Reqs key and the pre reqs key are in the same function because
// Sometimes the university of waterloo website will have Pre Requisites,
// or Co Requisites instead of Anti-Requisites, vice versa
//
// The function takes the cs: *CourseScrape parameter
func SetCourseAnti_Co_PreReqs(cs *CourseScrape) {
	if len(cs.Row) > 2 {
		// Start with anti reqs
		var splitBy, key string = "Antireq: ", "anti_reqs"

		// If it shows coreqs instead of antireqs, change the names
		if strings.Contains(cs.Row[2], "Coeq: ") {
			splitBy, key = "Coreq: ", "co_reqs"

			// If it shows prereqs instead of antireqs, change the names
		} else if strings.Contains(cs.Row[2], "Prereq: ") {
			splitBy, key = "Prereq: ", "pre_reqs"
		}

		// Split the string
		var split []string = strings.Split(cs.Row[2], splitBy)
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
// of having to call if strings.Contains() a bunch of times
//
// The function takes the cs: *CourseScrape parameter
func (cs *CourseScrape) IndexScrapeResult() {
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
		// Create a CourseScrap object
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
		cs.Row = strings.Split(splitTable[i], ">")[1:]

		// Check Data length
		if len(cs.Row[0]) > 1 {
			// Check if the splitTable contains a note about the course
			if strings.Contains(splitTable[i], "[Note: ") {
				SetCourseNote(cs, splitTable[i])
			} else {
				cs.Index++
				// Break the loop if the index is greater than 8
				if cs.Index > 8 {
					break
				}
				// Index the scrape result
				cs.IndexScrapeResult()
			}
		}
	}
	// Return the course id and the course info map (result)
	return cs.Result["id"], cs.Result
}

// The ScrapeCourseData() function is the main course scraper function
// This is because it scrapes all the course information and appends
// it to a map
//
// The function takes the client: *fasthttp.Client parameter to send http requests
//
// The function returns the course title string, the course data result map
// and the http request error
func ScrapeCourseData(client *fasthttp.Client, course string) (*map[string]map[string]string, error) {
	// Utilize the HttpRequest struct to easily send an http request
	var _Req *http.HttpRequest = &http.HttpRequest{
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
	if err != nil || resp.StatusCode() != 200 {
		return &result, err
	}

	// Define Variables
	// body: string -> The http response body
	// courseTables: []string -> The tables with each course program data
	// courseTitle: string -> The courses title (ex: CS -> computerscience)
	// waitGroup: sync.WaitGroup -> waitgroup for the scraping goroutines
	var (
		body         string          = string(resp.Body())
		courseTables []string        = strings.Split(body, "<div class=\"divTable\">")[1:]
		waitGroup    *sync.WaitGroup = &sync.WaitGroup{}
	)

	// Iterate over the html tables
	for _, table := range courseTables {
		waitGroup.Add(1)
		// Goroutine to scrape data
		go func(table *string) {
			defer waitGroup.Done()

			// Scrape the course data
			var courseID, courseData = _ScrapeCourseData(table)

			// Append the course data to the result map
			result[courseID] = courseData
		}(&table)
	}
	// Wait for all scraping goroutines to finish
	waitGroup.Wait()

	// Return the course title and it's result map containing
	// all the courses information
	return &result, nil
}
