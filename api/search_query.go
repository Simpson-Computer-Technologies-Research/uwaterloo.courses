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

// The GetSubjectNames() function returns the
// subject names from the global.SubjectNames map
// as a slice
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

		// Iterate using the smallest key length
		for i := 0; i < len(smallestKey); i++ {
			var (
				tempIndex    float64 = 1.0
				containCheck string  = ""
			)

			// If the keys equal the same
			if queryBytes[i] == subjectName[i] {
				largestResult += float64(queryBytes[i])
			} else {
				largestResult -= float64(int(queryBytes[i]) / len(largestKey))
			}

			// Iterate over the smallest key
			for j := 0; j < len(smallestKey); j++ {
				// Add the letter to the contain check string
				containCheck += string(string(queryBytes[j]))

				// Check if the subjectName contains the containCheck
				if strings.Contains(subjectName, containCheck) {
					// Make sure the length of the contain check
					// is greater than 2, or else you'll use single letters
					if len(containCheck) > 2 {
						largestResult += float64(
							float64(queryBytes[j]) / float64(len(containCheck)))
					}
				} else {
					// Reset the contain check
					containCheck = ""
				}
				// Get the distance the same letters are from eachother
				// using the tempIndex
				tempIndex++
				if subjectName[i] == smallestKey[j] {
					largestResult += float64(
						tempIndex / float64(len(smallestKey)*len(largestKey)))
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
	fmt.Printf(" >> Query: BestMatch (%v) (%s) (%v) \n\n",
		BestMatchValue, BestMatch, float64(370-(len(BestMatch)/2)))

	// Make sure best match is valid/accurate
	if BestMatchValue > float64(370-(len(BestMatch)/2)) {
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
