package api

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/realTristan/The_University_of_Waterloo/global"
)

// The CleanQuery() function removes all spaces from the query
// and removes all non alphabetic characters
func CleanQuery(query string) string {
	var res string = ""
	for i := 0; i < len(query); i++ {
		// Check if the character at the index is a letter
		if unicode.IsLetter(rune(query[i])) {
			// Append the letter to the result string
			res += string(query[i])
		}
	}
	// Return the res string in lowercase
	return strings.ToLower(res)
}

// The SearchQuery() function uses the cleaned query (ex: computerscience)
// and find the best match using it against the global.SubjectNames map
// It returns the best subject code match (ex: CS)
func SearchQuery(query string) string {
	// Clean the query
	query = CleanQuery(query)

	// Define Variables
	// bestMatchValue: int -> Track the highest value for character matching
	// bestMatch: string -> the best subject code for the query
	var (
		bestMatchValue int = -1
		bestMatch      string
	)

	// Iterate over the subject names map
	for subjectName, subjectCode := range global.SubjectNames {
		// Add the subject code to the result map
		var count int = 0

		// Iterate over the query characters
		for i := 0; i < len(query); i++ {
			// Check length so we don't get an error
			if i < len(subjectName) {
				// Check if the characters at the indexes are the same
				if subjectName[i] == query[i] {
					count += 1 * (len(query) / len(subjectName))
				}
			}
			// Query vs SubjectName length
			if i+2 > len(query)/((len(subjectName)/len(query))+1) {
				break
			}
		}
		// Check if the current subject is more accurate
		// than the previous ones
		if count > bestMatchValue {
			bestMatchValue = count
			bestMatch = subjectCode
		}
	}
	// Print the best match data for testing
	fmt.Printf(
		"Query BestMatchValue: %d\nQueryBestMatch %s\n", bestMatchValue, bestMatch)

	// Check if the bestMatch is valid
	if bestMatchValue > 4 {
		return bestMatch
	}
	return ""
}

// The QueryHandler() function handles the search query and whether
// to use the native course arg or the query arg
//
// It'll also check for special searches for example: @code:
// will search for a specific subject code instead of for example:
// searching "computer science"
func QueryHandler(r *http.Request) string {
	// Define Variables
	// course: string -> the course code arg
	// query: string -> the course search query arg
	var (
		course string = string(r.URL.Query().Get("course"))
		query  string = string(r.URL.Query().Get("q"))
	)

	// Check if the query contains a special search
	if len(course) == 0 && len(query) > 0 {
		// Check if the user is searching for a specific subject code
		if strings.Contains(query, "@code") {
			return strings.ToUpper(
				CleanQuery(strings.Split(query, "@code")[1]))
		}
		// If using a search query (ex: computerscience) then match the query
		// to a subject code
		return SearchQuery(query)
	}
	// Return the course arg
	return course
}
