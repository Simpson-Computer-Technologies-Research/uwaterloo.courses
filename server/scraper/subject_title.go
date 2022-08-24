package scraper

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/realTristan/uwaterloo.courses/server/requests"
)

// The CleanSubjectTitle() function will remove all spaces and
// new lines from the h2 header
//
// The function returns the cleaned title string
func CleanSubjectTitle(title string) string {
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

// The ScrapeSubjectTitle() function will return the title of the subject at
// the top of the "https://classes.uwaterloo.ca/uwpcshtm.html" website
//
// This title will be used for search indexing thus it needs to be cleaned using
// the CleanSubjectTitle() function
//
// The function uses the scraped title to return the cleaned subject title string
func _ScrapeSubjectTitle(body *string) string {
	// First header split
	var temp []string = strings.Split(*body, "<h2 class=\"subject\">")
	// Check length of first split
	if len(temp) > 1 {
		// Second header split
		var _temp []string = strings.Split(temp[1], "</h2>")

		// Check the length of the secind split
		if len(_temp) > 0 {
			// Clean the subject title and return it
			// Example: C O M P U T E R -> computer
			return CleanSubjectTitle(_temp[0])
		}
	}
	// Return an empty string
	return ""
}

// The ScrapeSubjectTitle() is used to send an http request to the
// official university of waterloo website and return the scraped
// subject codes using the _ScrapeSubjectTitle(&response_body) function
func ScrapeSubjectTitle(course string) string {
	var (
		// Create the request object used for sending the http request
		_Req *requests.HttpRequest = &requests.HttpRequest{
			Url:    fmt.Sprintf("https://ucalendar.uwaterloo.ca/2223/COURSE/course-%s.html", course),
			Method: "GET",
			Headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
			},
		}
		// resp, err -> Http response and errors
		resp, err = _Req.Send()
	)
	// Handle any errors
	if err != nil {
		fmt.Printf(" >> Scrape Subject Title Error: %e", err)
		return ""
	}
	// Convert the response body to a string
	var body string = string(resp.Body())

	// Return the scraped subject titles
	return _ScrapeSubjectTitle(&body)
}
