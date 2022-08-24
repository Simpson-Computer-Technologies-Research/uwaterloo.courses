package api

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/realTristan/uwaterloo.courses/server/global"
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
	return res
}

// The GetLargest() function returns the
// largest byte out of the two
func GetLargest(a []byte, b []byte) []byte {
	if len(a) > len(b) {
		return a
	}
	return b
}

// The GetSmallest() function returns the
// smallest byte out of the two
func GetSmallest(a []byte, b []byte) []byte {
	if len(a) > len(b) {
		return b
	}
	return a
}

// The GetBestMatch() function uses the cleaned query (ex: computerscience)
// to find the best match using the global.SubjectNames map
// It returns the best subject code match (ex: CS)
//
// I tried my best to use non constants when adding to the
// best match value. The searching is actually pretty accurate
// For Example, I was able to match "cheistyr" with "chemistry"
//
// I set everything to float64 so the decimals can play a role
// in micro differences
func GetBestMatch(query string) string {
	// Define the bestmatch beginning values
	var (
		bestMatch      string  = ""
		bestMatchValue float64 = -1.0

		// Query converted to bytes with all spaces removed
		queryBytes []byte = []byte(CleanQuery(query))
	)

	// Iterate over the subject names
	for subjectName := range global.SubjectNames {
		var (
			// Result value
			resVal float64 = 0.0

			// Get the largest / smallest keys
			largestKey  []byte = GetLargest([]byte(subjectName), queryBytes)
			smallestKey []byte = GetSmallest([]byte(subjectName), queryBytes)
		)

		// Iterate using the smallest key length
		for i := 0; i < len(smallestKey); i++ {
			var (
				tmpIndx float64 = 1.0
				substr  string  = ""
			)

			// If the keys equal the same
			if queryBytes[i] == subjectName[i] {
				resVal += float64(queryBytes[i])
			} else {
				resVal -= float64(queryBytes[i]) / float64(len(largestKey))
			}

			// Iterate over the smallest key
			for j := 0; j < len(smallestKey); j++ {
				// Add the letter to the substr
				substr += string(queryBytes[j])

				// Check if the subjectName contains the substr
				if strings.Contains(subjectName, substr) {
					// Make sure the length of the contain check
					// is greater than 2, or else you'll use single letters
					if len(substr) > 2 {
						resVal += float64(queryBytes[j]) / float64(len(substr))
					}
				} else {
					// Reset the contain check
					substr = ""
				}
				// Get the distance the same letters are from eachother
				// using the tempIndex
				tmpIndx++
				if subjectName[i] == smallestKey[j] {
					resVal += tmpIndx / float64(len(smallestKey)*len(largestKey))
				}
			}
			// Check if smallest key contains the subject name letter
			if !strings.Contains(string(smallestKey), string(subjectName[i])) {
				resVal -= float64(len(smallestKey))
			}
		}
		// Check if resval is greater than
		// the previous bestmatchvalues
		if resVal > bestMatchValue {
			bestMatchValue = resVal
			bestMatch = subjectName
		}
	}
	// Make sure best match is valid/accurate
	if bestMatchValue > float64(370-(len(bestMatch)/2)) {
		// Return the best match subject code
		return global.SubjectNames[bestMatch]
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
	var query string = strings.ToLower(r.URL.Query().Get("q"))

	// Check if the user is searching for a specific subject code
	if strings.Contains(query, "@code") {
		return strings.ToUpper(
			CleanQuery(strings.Split(query, "@code")[1]))
	}
	// If using a search query (ex: computerscience) then match the query
	// to a subject code
	return GetBestMatch(query)
}
