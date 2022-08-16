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

// Get the largest byte out of the two
func GetLargest(a []byte, b []byte) []byte {
	if len(a) > len(b) {
		return a
	}
	return b
}

// Get the smallest byte out of the two
func GetSmallest(a []byte, b []byte) []byte {
	if len(a) > len(b) {
		return b
	}
	return a
}

// Get the subjectnames as a string golang
func GetSubjectNames() []string {
	var result []string = []string{}
	for k := range global.SubjectNames {
		result = append(result, k)
	}
	return result
}

// The GetBestMatch() function uses the cleaned query (ex: computerscience)
// and find the best match using it against the global.SubjectNames map
// It returns the best subject code match (ex: CS)

func GetBestMatch(query string) string {
	// Define the bestmatch beginning values
	var (
		BestMatch      string  = ""
		BestMatchValue float64 = -1.0
	)

	// Iterate over the subject names
	for subjectName := range global.SubjectNames {
		var (
			// Largest result and the search query as bytes + replace all spaces
			largestResult float64 = 0.0
			queryBytes    []byte  = []byte(strings.ReplaceAll(query, " ", ""))

			// Get the largest / smallest keys
			largestKey  []byte = GetLargest([]byte(subjectName), queryBytes)
			smallestKey []byte = GetSmallest([]byte(subjectName), queryBytes)
		)

		// Iterations
		for i := 0; i < len(smallestKey); i++ {
			var tempIndex float64 = 1.0

			// If the keys equal the same
			if queryBytes[i] == subjectName[i] {
				largestResult += float64(queryBytes[i])
			} else {
				largestResult -= float64(int(queryBytes[i]) / len(largestKey))
			}
			// Iterate over the smallest key
			for j := 0; j < len(smallestKey); j++ {
				tempIndex++
				// Get the distance the same letters are from eachother
				// using the tempIndex
				if subjectName[i] == smallestKey[j] {
					largestResult += float64(tempIndex / float64(len(smallestKey)*len(largestKey)))
				}
			}
			// Check if smallest key contains the subject name letter
			if !strings.Contains(string(smallestKey), string(subjectName[i])) {
				largestResult -= float64(len(smallestKey))
			}
		}

		// Check if bestmatchvalue is greater than
		// the previous bestmatchvalues
		if largestResult > BestMatchValue {
			BestMatchValue = largestResult
			BestMatch = subjectName
		}
	}
	// Log the best match
	fmt.Printf(" >> Query: BestMatch (%v) (%s)\n\n", BestMatchValue, BestMatch)

	// Make sure best match is valid/accurate
	if BestMatchValue > 370 {
		// Return the best match subject code
		return global.SubjectNames[BestMatch]
	}
	// Return None
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

	if len(course) == 0 && len(query) > 0 {
		// Check if the user is searching for a specific subject code
		if strings.Contains(query, "@code") {
			return strings.ToUpper(
				CleanQuery(strings.Split(query, "@code")[1]))
		}
		// If using a search query (ex: computerscience) then match the query
		// to a subject code
		return GetBestMatch(query)
	}
	// Return the course arg
	return course
}
