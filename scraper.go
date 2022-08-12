package main

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

// The _ScrapSubjectCodes() function will return a slice containing all
// the course codes from the provided html
// The html is from: https://classes.uwaterloo.ca/uwpcshtm.html
func _ScrapeSubjectCodes(html string) []string {
	// Declare Variables
	// res: []string -> result slice holding the subject codes
	// tableIndex: int -> used to only append the subject codes to res
	var (
		res        []string = []string{}
		tableIndex int      = 1
	)
	html = strings.Split(html, "</table>")[1]

	// Iterate over the split strings
	for i, p := range strings.Split(html, "<td>") {
		// Get every 7th table value (the Subject table)
		if i == tableIndex {
			// Increase tableIndex by 7
			tableIndex += 7
			// Split by closing tag
			var s []string = strings.Split(p, "</td>")
			// If the result slice doesn't contains the subject
			if !SliceContains(res, s[0]) {
				// Append the subject to the result slice
				res = append(res, s[0])
			}
		}
	}
	// Return the result slice
	return res
}

// The ScrapSubjectCodes() function will utilizes the _ScrapeSubjectCodes()
// function to webscrape all the subject codes on the
// "https://classes.uwaterloo.ca/uwpcshtm.html" website
//
// The Subject Codes will be used to get information about the courses
// from the "https://ucalendar.uwaterloo.ca/2021/COURSE/course-CS.html"
// website.
func ScrapeSubjectCodes(client *fasthttp.Client) ([]string, error) {
	// Utilize the _Request struct to easily send an http request
	var _Req *_Request = &_Request{
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

	// Scrape the subject codes using the response.body()
	// Then return the codes alongside no error
	return _ScrapeSubjectCodes(string(resp.Body())), nil
}

// The IndexCourseInfoScrapeResult() will index all the course information into a map
// Notes: yes i know this code is nasty but I didn't know what else to do.
// PLEASE give me suggestions!
func IndexCourseInfoScrapeResult(index int, data []string, result map[string]string) map[string]string {
	if index == 1 {
		// Set the "Course Info" key, this contains the course name, catalog num, components and unit
		result["Course Info"] = data[0]
		//
	} else if index == 2 {
		// Set the "Course ID" key, this is the unique int id of the course
		result["Course ID"] = strings.Split(data[1], "Course ID: ")[1]
		//
	} else if index == 3 {
		// Set the "Course Name" key, this is the courses name
		result["Course Name"] = data[2]
		//
	} else if index == 4 {
		// Set the "Course Description" key, this is a description of the course
		result["Course Description"] = data[1]
		//
	} else if index == 6 {
		// Set the "Pre-Reqs" key, these are all the required requisites
		result["Pre-Reqs"] = strings.Split(data[2], "Prereq: ")[1]
		//
	} else if index == 7 {
		// Set the "Anti-Reqs" key, these are the requisites you can't have
		result["Anti-Reqs"] = strings.Split(data[2], "Antireq: ")[1]
		//
	} else if index == 8 {
		// Set the "Other" key, this key is usually an "Online Only" url
		result["Other"] = strings.Split(data[2], "<a href=")[1]
	}
	// Return the result map
	return result
}

// The _ScrapeCourseInfo() function will create a result map
// that stores the course info.  The course info map holds the
// course id, name, description, pre-reqs, anti-reqs, etc.
func _ScrapeCourseInfo(table string) (string, map[string]string) {
	// Define Variables
	// result: map[string]string -> The result map that holds the course info
	var result map[string]string = make(map[string]string)

	// Create a seperate index variable for the IndexCourseInfoScrapeResult() function
	// Split the table into the segments that contain the course info
	var (
		index      int      = 0
		splitTable []string = strings.Split(table, "</")[1:]
	)
	// Iterate through each segment
	for i := 0; i < len(splitTable); i++ {
		// Split the segment by >
		var data []string = strings.Split(splitTable[i], ">")[1:]
		// Increase index variable by 1
		if len(data[0]) > 1 {
			index++
			// Append the result
			result = IndexCourseInfoScrapeResult(index, data, result)
		}
	}
	// Return the course id and the course info map (result)
	return result["Course ID"], result
}

// The ScrapeCourseInfo() function
func ScrapeCourseInfo(client *fasthttp.Client, course string) map[string]map[string]string {
	// Utilize the _Request struct to easily send an http request
	var _Req *_Request = &_Request{
		Client: client,
		Url:    fmt.Sprintf("https://ucalendar.uwaterloo.ca/2021/COURSE/course-%s.html", course),
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
		},
	}
	// Send Http Request
	var resp, _ = _Req.Send()

	// Define Variables
	// courseTables: []string -> The tables with each course program info
	// result: map[string]map[string]string -> The result map that holds all the course data
	var (
		courseTables []string = strings.Split(string(resp.Body()), "<div class=\"divTable\">")[1:]
		result                = make(map[string]map[string]string)
	)

	// Iterate over the html tables
	for _, table := range courseTables {
		var courseID, courseInfo = _ScrapeCourseInfo(table)
		result[courseID] = courseInfo
	}
	return result
}
