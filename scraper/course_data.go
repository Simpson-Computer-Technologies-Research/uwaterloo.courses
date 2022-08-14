package scraper

// Import packages
import (
	"fmt"
	"strings"
	"sync"
	"time"

	http "github.com/realTristan/The_University_of_Waterloo/http"
	"github.com/valyala/fasthttp"
)

// The Course Scrape struct holds three keys
/* - Index: int -> used to select which Result key to use 		*/
/* - Row: []string -> the row with data							*/
/* - Result: map[string]string -> the course data result map	*/
type CourseScrape struct {
	Row    []string
	Result map[string]string
	HTML   string
}

// The Course Scrape struct holds four keys
/* - ResultHTML: string -> The scraped data html result 				 */
/* - ResultSlice: string -> The scraped data result slice 				 */
/* - Mutex: *sync.RWMutex -> The Mutex Lock for prevent data overwrites  */
/* - WaitGroup: *sync.WaitGroup -> The Wait Group for goroutines         */
type ScrapeResult struct {
	ResultHTML  string
	ResultSlice []map[string]string
	Mutex       *sync.RWMutex
	WaitGroup   *sync.WaitGroup
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
// The function takes the cs: *CourseScrape parameter
func (cs *CourseScrape) SetCourseInfo() {
	if len(cs.Row) > 0 {
		// Index the course info
		var title, comps, unit = IndexCourseInfo(cs.Row[0])

		// Append values to the result map
		cs.Result["title"] = title
		cs.Result["components"] = comps
		cs.Result["unit"] = unit

		// Append values to the html
		cs.AppendHTML(title, "")
		cs.AppendHTML("Components", comps)
		cs.AppendHTML("Unit", unit)
	}
}

// The SetCourseId() function will set the CourseScrape
// Result map id. The Course id is the key used to categorize all the data
// from each course. Example: id: {data: values}
//
// The function takes the cs: *CourseScrape parameter
func (cs *CourseScrape) SetCourseId() {
	// Data Length is greater than one
	if len(cs.Row) > 1 {
		// Split the string by "Course ID: " to get the exact id
		var split []string = strings.Split(cs.Row[1], "Course ID: ")
		// Make sure the split length is greater than one
		if len(split) > 1 {
			// Set the result map value
			cs.Result["id"] = split[1]
			cs.AppendHTML("ID", split[1])
		}
	}
}

// The SetCourseName() function will set the CourseScrape
// Result map name -> Example result["name"] = "Intro to Computer Science"
//
// The function takes the cs: *CourseScrape parameter
func (cs *CourseScrape) SetCourseName() {
	if len(cs.Row) > 2 {
		cs.Result["name"] = cs.Row[2]
		cs.AppendHTML("Name", cs.Row[2])
	}
}

// The SetCourseDescription() function will set the CourseScrape
// Result map description -> Example result["desc"] = "Learn about..."
//
// The function takes the cs: *CourseScrape parameter
func (cs *CourseScrape) SetCourseDescription() {
	if len(cs.Row) > 1 {
		cs.Result["desc"] = cs.Row[1]
		cs.AppendHTML("Description", cs.Row[1])
	}
}

// The SetCourseNote() function will set the CourseScrape
// Result map note -> Example result["note"] = "Only available for..."
//
// The function takes the cs: *CourseScrape parameter
// and the data string parameter which is the html row of the note
func (cs *CourseScrape) SetCourseNote(data string) {
	// Split the string to get note content
	var split []string = strings.Split(data, "[Note: ")
	if len(split) > 1 {
		// Set the note in the result map
		cs.Result["note"] = "[" + split[1]
		cs.AppendHTML("Note", "["+split[1])
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
func (cs *CourseScrape) SetCourseAnti_Co_PreReqs() {
	if len(cs.Row) > 2 {
		// Start with anti reqs
		var splitBy, key, name string = "Antireq: ", "anti_reqs", "Anti Requisites"

		// If it shows coreqs instead of antireqs, change the names
		if strings.Contains(cs.Row[2], "Coeq: ") {
			splitBy, key, name = "Coreq: ", "co_reqs", "Co Requisites"

			// If it shows prereqs instead of antireqs, change the names
		} else if strings.Contains(cs.Row[2], "Prereq: ") {
			splitBy, key, name = "Prereq: ", "pre_reqs", "Pre Requisites"
		}

		// Split the string
		var split []string = strings.Split(cs.Row[2], splitBy)
		if len(split) > 1 {
			cs.Result[key] = split[1]
			cs.AppendHTML(name, split[1])
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
func (cs *CourseScrape) IndexScrapeResult(index int) {
	// The map for the CourseScrape.Index
	var indexMap map[int]func() = map[int]func(){
		1: cs.SetCourseInfo,            // title, components, unit
		2: cs.SetCourseId,              // id
		3: cs.SetCourseName,            // name
		4: cs.SetCourseDescription,     // desc
		5: func() {},                   // [empty]
		6: cs.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		7: cs.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
		8: cs.SetCourseAnti_Co_PreReqs, // anti_reqs or co_reqs or pre_reqs
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
func _ScrapeCourseData_(table *string) (map[string]string, string) {
	// Define Variables
	var (
		// Create a CourseScrap object
		cs *CourseScrape = &CourseScrape{
			Result: make(map[string]string),
			HTML:   "",
		}
		// Split the table into the segments that contain the course info
		splitTable []string = strings.Split(*table, "</")[1:]

		// Track table index
		index int = 0
	)

	// Iterate through the split table
	for i := 0; i < len(splitTable); i++ {
		// Split the splitted table by >
		cs.Row = strings.Split(splitTable[i], ">")[1:]

		// Check Data length
		if len(cs.Row[0]) > 1 {
			// Check if the splitTable contains a note about the course
			if strings.Contains(splitTable[i], "[Note: ") {
				cs.SetCourseNote(splitTable[i])
			} else {
				index++
				// Break the loop if the index is greater than 8
				if index > 8 {
					break
				}
				// Index the scrape result
				cs.IndexScrapeResult(index)
			}
		}
	}
	// Return the course id and the course info map (result)
	return cs.Result, cs.WrapHTML()
}

// The _ScrapeCourseData_() function uses the ScrapeResult
// object to prevent deadlocking, data overwriting and
// uses the waitgroup for waiting until the goroutine finishes
//
// The function is called through a goroutine to maximize
// scrape speed
func (sr *ScrapeResult) _ScrapeCourseData(t string) {
	// Finish goroutine wait once function returns
	defer sr.WaitGroup.Done()

	// Lock the Mutex
	sr.Mutex.Lock()

	// Unlock the mutex once the function returns
	defer sr.Mutex.Unlock()

	// Scrape course data
	var courseData, htmlData = _ScrapeCourseData_(&t)

	// Append the course data to the result map
	sr.ResultSlice = append(sr.ResultSlice, courseData)

	// Append the html data to the html result
	sr.ResultHTML += fmt.Sprintf("<br><br>%s", htmlData)
}

// The ScrapeCourseData() function is the main course scraper function
// This is because it scrapes all the course information and appends
// it to a map
//
// The function takes the client: *fasthttp.Client parameter to send http requests
//
// The function returns the course title string, the course data result map
// the result html and the http request error
func ScrapeCourseData(client *fasthttp.Client, course string) (*[]map[string]string, string, error) {
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
	// scrapeResult: *ScrapeResult -> Holds all the scrape data, the mutex lock and the waitgroup
	var (
		resp, err                  = _Req.Send()
		scrapeResult *ScrapeResult = &ScrapeResult{
			ResultSlice: []map[string]string{},
			Mutex:       &sync.RWMutex{},
			WaitGroup:   &sync.WaitGroup{},
			ResultHTML:  "",
		}
	)

	// Handle response error
	if err != nil || resp.StatusCode() != 200 {
		return &scrapeResult.ResultSlice, "", err
	}

	// Define Variables
	var (
		// Track how long it takes to scrape data
		scrapeStartTime time.Time = time.Now()
		// The response body string
		body string = string(resp.Body())
		// The courseTables slice
		courseTables []string = strings.Split(body, "<div class=\"divTable\">")[1:]
	)

	// Iterate over the html tables
	for _, table := range courseTables {
		scrapeResult.WaitGroup.Add(1)
		go scrapeResult._ScrapeCourseData(table)
	}

	// Wait for scraping to finish
	scrapeResult.WaitGroup.Wait()

	// Log the time it took to scrape the course data
	// It usually takes around 1-20ms
	fmt.Printf(" [LOG] Scraped Course Data [%v]\n", time.Since(scrapeStartTime))

	// Return the result map containing all the
	// course information, the html result data and the
	// http request error
	return &scrapeResult.ResultSlice, scrapeResult.ResultHTML, nil
}
