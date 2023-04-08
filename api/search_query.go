package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
	"unicode"

	"github.com/realTristan/uwaterloo.courses/global"
)

// The CleanQuery() function removes all spaces from the query
// and removes all non alphabetic characters
func CleanQuery(query []byte) []byte {
	var res []byte
	for i := 0; i < len(query); i++ {
		// Check if the character at the index is a letter
		if unicode.IsLetter(rune(query[i])) {
			// Append the letter to the result
			res = append(res, query[i])
		}
	}
	// Return the res in lowercase
	return bytes.ToLower(res)
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
func GetBestMatch(query []byte) []byte {
	// Define the bestmatch beginning values
	var (
		// Track get best match time
		startTime time.Time = time.Now()

		// Best match values
		bestMatch      []byte
		bestMatchValue float64 = -1.0
	)

	// Iterate over the subject names
	for _subjectName := range global.SubjectNames {
		var (
			// Convert subject name to bytes
			subjectName []byte = []byte(_subjectName)

			// Result value
			resVal float64 = 0.0

			// Get the largest / smallest keys
			largestKey  []byte = GetLargest(subjectName, query)
			smallestKey []byte = GetSmallest(subjectName, query)
		)

		// Iterate using the smallest key length
		for i := 0; i < len(smallestKey); i++ {
			var (
				tmpIndx float64 = 1.0
				substr  []byte
			)

			// If the keys equal the same
			if query[i] == subjectName[i] {
				resVal += float64(query[i])
			} else {
				resVal -= float64(query[i]) / float64(len(largestKey))
			}

			// Iterate over the smallest key
			for j := 0; j < len(smallestKey); j++ {
				// Add the letter to the substr
				substr = append(substr, query[j])

				// Check if the subjectName contains the substr
				if bytes.Contains(subjectName, substr) {
					// Make sure the length of the contain check
					// is greater than 2, or else you'll use single letters
					if len(substr) > 2 {
						resVal += float64(query[j]) / float64(len(substr))
					}
				} else {
					// Reset the contain check
					substr = []byte{}
				}
				// Get the distance the same letters are from eachother
				// using the tempIndex
				tmpIndx++
				if subjectName[i] == smallestKey[j] {
					resVal += tmpIndx / float64(len(smallestKey)*len(largestKey))
				}
			}
		}
		// Check if resval is greater than
		// the previous bestmatchvalues
		if resVal > bestMatchValue {
			bestMatchValue = resVal
			bestMatch = subjectName
		}
	}
	// Print the query result
	fmt.Printf("\n >> Best Match Query: (%s) (%f) (%v)\n",
		bestMatch, bestMatchValue, time.Since(startTime))

	// Make sure best match is valid/accurate
	if bestMatchValue > float64(370-(len(bestMatch)/2)) {
		// Return the best match subject code
		return []byte(global.SubjectNames[string(bestMatch)])
	}
	// Return None
	return []byte{}
}

// The QueryHandler() function handles the search query and whether
// to use the native course arg or the query arg
//
// It'll also check for special searches for example: @code:
// will search for a specific subject code instead of for example:
// searching "computer science"
func QueryHandler(r *http.Request) []byte {
	// Define Variables
	// query: string -> the course search query arg
	// codeByte: []byte -> the @code bytes
	var (
		query     []byte = bytes.ToLower([]byte(r.URL.Query().Get("q")))
		codeBytes []byte = []byte("@code")
	)

	// Check if the user is searching for a specific subject code
	if bytes.Contains(query, codeBytes) {
		return bytes.ToUpper(
			CleanQuery(bytes.Split(query, codeBytes)[1]))
	}
	// If using a search query (ex: computerscience) then match the query
	// to a subject code
	return GetBestMatch(CleanQuery(query))
}
