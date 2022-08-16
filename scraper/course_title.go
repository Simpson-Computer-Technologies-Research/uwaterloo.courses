package scraper

import (
	"strings"
	"unicode"
)

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

// The ScrapeCourseTitle() function will return the title of the course at
// the top of the "https://classes.uwaterloo.ca/uwpcshtm.html" website
//
// This title will be used for search indexing thus it needs to be cleaned using
// the CleanCourseTitle() function
//
// The function uses the scraped title to return the cleaned course title string
func ScrapeCourseTitle(body *string) string {
	// First header split
	var temp []string = strings.Split(*body, "<h2 class=\"subject\">")
	// Check length of first split
	if len(temp) > 1 {
		// Second header split
		var _temp []string = strings.Split(temp[1], "</h2>")

		// Check the length of the secind split
		if len(_temp) > 0 {
			// Clean the course title and return it
			// Example: C O M P U T E R -> computer
			return CleanCourseTitle(_temp[0])
		}
	}
	// Return an empty string
	return ""
}
