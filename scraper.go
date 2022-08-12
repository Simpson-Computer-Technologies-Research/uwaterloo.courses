package main

import (
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
