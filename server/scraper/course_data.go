package scraper

// Import packages
import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/realTristan/uwaterloo.courses/server/cache"
	"github.com/realTristan/uwaterloo.courses/server/global"
	"github.com/realTristan/uwaterloo.courses/server/requests"
	"github.com/valyala/fasthttp"
)

// The Course Scrape struct holds three keys
/* - Row: []string -> the row with data							*/
/* - Result: map[string]string -> the course data result map	*/
type ScrapeTable struct {
	Row    []string
	Result map[string]string
}

// The Course Scrape struct holds four keys
/* - ResultSlice: string -> The scraped data result slice 				 */
/* - Mutex: *sync.RWMutex -> The Mutex Lock for prevent data overwrites  */
/* - WaitGroup: *sync.WaitGroup -> The Wait Group for goroutines         */
type ScrapeResult struct {
	ResultSlice []map[string]string
	Mutex       *sync.RWMutex
	WaitGroup   *sync.WaitGroup
}

// The RefreshCache() function re-scrapes and re-sets all
// the keys in the redis cache database to the scrape result
func RefreshCache() {
	var (
		RequestClient *fasthttp.Client = &fasthttp.Client{}
		s3Result      []map[string]string
	)

	// Iterate over the subject codes
	for _, v := range global.SubjectCodes {
		s3Result = append(s3Result, ScrapeCourseData(RequestClient, v).ResultSlice...)
	}

	// Write to the json file
	var res, _ = json.Marshal(s3Result)
	os.WriteFile("./default_data.txt", res, 0644)
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
	var splitTitle []string = strings.Split(title, " ")
	if len(splitTitle) > 1 {
		var courseTitle string = fmt.Sprintf("%s %s", splitTitle[0], splitTitle[1])
		// Return the Course Title ex: CS 201
		// Return the Components (splitTitle[2]) ex: LAB,LEC,TST
		// Return the Unit (splitTitle[3]) ex: 0.50
		return courseTitle, splitTitle[2], splitTitle[3]
	}
	// Return Empty
	return "", "", ""
}

// The SetCourseInfo() function takes in the original course info
// and sets the title, components and unit using the IndexCourseInfo
// function. This makes it much easier to read on the front end
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) SetCourseInfo() {
	if len(st.Row) > 0 {
		// Index the course info
		st.Result["title"],
			st.Result["components"],
			st.Result["unit"] = IndexCourseInfo(st.Row[0])
	}
}

// The SetCourseId() function will set the ScrapeTable
// Result map id. The Course id is the key used to categorize all the data
// from each course. Example: id: {data: values}
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) SetCourseId() {
	// Data Length is greater than one
	if len(st.Row) > 1 {
		// Split the string by "Course ID: " to get the exact id
		var split []string = strings.Split(st.Row[1], "Course ID: ")
		// Make sure the split length is greater than one
		if len(split) > 1 {
			// Set the id key in the result map
			st.Result["id"] = split[1]
		}
	}
}

// The SetCourseName() function will set the ScrapeTable
// Result map name -> Example result["name"] = "Intro to Computer Science"
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) SetCourseName() {
	if len(st.Row) > 2 {
		// Set the name key in the result map
		st.Result["name"] = st.Row[2]
	}
}

// The SetCourseDescription() function will set the ScrapeTable
// Result map description -> Example result["desc"] = "Learn about..."
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) SetCourseDescription() {
	if len(st.Row) > 1 {
		// Set the description in the result map
		st.Result["description"] = st.Row[1]
	}
}

// The SetCourseNote() function will set the ScrapeTable
// Result map note -> Example result["note"] = "Only available for..."
//
// The function takes the cs: *ScrapeTable parameter
// and the data string parameter which is the html row of the note
func (st *ScrapeTable) SetCourseNote(data string) {
	// Split the string to get note content
	var split []string = strings.Split(data, "[Note: ")
	if len(split) > 1 {
		// Set the note in the result map
		st.Result["note"] = "[" + split[1]
	}
}

// The SetCourseAnti_Co_PreReqs() function will set the ScrapeTable
// Result map anti_reqs, co_reqs or pre_reqs key
//
// The Anti Reqs key, Co Reqs key and the pre reqs key are in the same function because
// Sometimes the university of waterloo website will have Pre Requisites,
// or Co Requisites instead of Anti-Requisites, vice versa
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) SetCourseAnti_Co_PreReqs() {
	if len(st.Row) > 2 {
		// Start with anti reqs
		var splitBy, name string = "Antireq: ", "Anti Requisites"

		// If it shows coreqs instead of antireqs, change the names
		if strings.Contains(st.Row[2], "Coeq: ") {
			splitBy, name = "Coreq: ", "Co Requisites"

			// If it shows prereqs instead of antireqs, change the names
		} else if strings.Contains(st.Row[2], "Prereq: ") {
			splitBy, name = "Prereq: ", "Pre Requisites"
		}

		// Split the string
		var split []string = strings.Split(st.Row[2], splitBy)
		if len(split) > 1 {
			// Set the key in the result map
			// Append the key to the html result
			st.Result[strings.ToLower(name)] = split[1]
		}
	}
}

// The IndexScrapeResult() uses a map to categorize all the functions
// that will be used for indexing the scrap result
//
// I decided to go with this way because it was cleaner than having
// a bunch of if and else if statements
//
// I also decided to use an int index for categorizing everything instead
// of having to call if strings.Contains() a bunch of times
//
// The function takes the cs: *ScrapeTable parameter
func (st *ScrapeTable) IndexScrapeResult(index int) {
	// The map for the table scrape index
	var indexMap map[int]func() = map[int]func(){
		1: st.SetCourseInfo,            // title, components, unit
		2: st.SetCourseId,              // id
		3: st.SetCourseName,            // name
		4: st.SetCourseDescription,     // desc
		5: func() {},                   // [empty]
		6: st.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		7: st.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		8: st.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
	}
	// Call the function
	indexMap[index]()
}

// The _ScrapeCourseData() function will create a result map
// that stores the course data. The course data map holds the
// course id, name, description, pre-reqs, anti-reqs, etc.
//
// The function takes the table: *string parameter
//
// The function returns the result map[string]string
func _ScrapeCourseData_(st *ScrapeTable, table string) map[string]string {
	// Define Variables
	// splitTable: []string -> The table into the segments that contain the course info
	// tableIndex: int -> Track table index
	var (
		splitTable []string = strings.Split(table, "</")[1:]
		tableIndex int      = 0
	)

	// Iterate through the split table
	for i := 0; i < len(splitTable); i++ {
		// Split the splitted table by >
		st.Row = strings.Split(splitTable[i], ">")[1:]

		// Check Data length
		if len(st.Row[0]) > 1 {
			// Check if the splitTable contains a note about the course
			if strings.Contains(splitTable[i], "[Note: ") {
				st.SetCourseNote(splitTable[i])
			} else {
				tableIndex++
				// Break the loop if the index is greater than 8
				if tableIndex > 8 {
					break
				}
				// Index the scrape result
				st.IndexScrapeResult(tableIndex)
			}
		}
	}

	// Return the course id and the course info map (result)
	return st.Result
}

// The _ScrapeCourseData_() function uses the ScrapeResult
// object to prevent deadlocking, data overwriting and
// uses the waitgroup for waiting until the goroutine finishes
//
// The function is called through a goroutine to maximize
// scrape speed
func (sr *ScrapeResult) _ScrapeCourseData(table string) {
	// Finish goroutine wait once function returns
	defer sr.WaitGroup.Done()

	// Lock the Mutex
	// Then Unlock it once the function returns
	sr.Mutex.Lock()
	defer sr.Mutex.Unlock()

	// Scrape course data, pass the ScrapeTable object
	var courseData = _ScrapeCourseData_(&ScrapeTable{
		Result: make(map[string]string),
	}, table)

	// Append course data to the cache
	cache.Set(courseData)

	// Append the course data to the result map
	sr.ResultSlice = append(sr.ResultSlice, courseData)
}

// The ScrapeCourseData() function is the main course scraper function
// This is because it scrapes all the course information and appends
// it to a map
//
// The function takes the client: *fasthttp.Client parameter to send http requests
//
// The function returns the course data result slice,
// the result html and the http request error
func ScrapeCourseData(client *fasthttp.Client, subjectCode string) *ScrapeResult {
	subjectCode = strings.ToUpper(subjectCode)

	// Utilize the HttpRequest struct to easily send an http request
	var _Req *requests.HttpRequest = &requests.HttpRequest{
		Client: client,
		Url:    fmt.Sprintf("https://ucalendar.uwaterloo.ca/2223/COURSE/course-%s.html", subjectCode),
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
		},
	}
	// Define Variables
	// resp, err -> request response and error
	// scrapeResult: *ScrapeResult -> Holds all the scrape data,
	// the mutex lock and the waitgroup
	var (
		resp, err                  = _Req.Send()
		scrapeResult *ScrapeResult = &ScrapeResult{
			ResultSlice: []map[string]string{},
			Mutex:       &sync.RWMutex{},
			WaitGroup:   &sync.WaitGroup{},
		}
	)
	// Handle response error
	if err != nil || resp.StatusCode() != 200 {
		return scrapeResult
	}

	// Define Variables
	var (
		// Track how long it takes to scrape data
		scrapeStartTime time.Time = time.Now()
		// The response body string
		body string = string(resp.Body())
		// The course tables slice
		courseTables []string = strings.Split(body, "<div class=\"divTable\">")[1:]
	)

	// Iterate over the html tables
	for i := 0; i < len(courseTables); i++ {
		scrapeResult.WaitGroup.Add(1)
		go scrapeResult._ScrapeCourseData(courseTables[i])
	}

	// Wait for scraping to finish
	scrapeResult.WaitGroup.Wait()

	// Log the time it took to scrape the course data
	// It usually takes around 500µs -> 3ms
	fmt.Printf(" >> Scraped Course Data [%v]\n\n", time.Since(scrapeStartTime))

	// Return the scrape result struct
	return scrapeResult
}
