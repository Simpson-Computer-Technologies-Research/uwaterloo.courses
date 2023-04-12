package api

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/realTristan/uwaterloo.courses/global"
)

// The CleanQuery() function removes all spaces from the query
// and removes all non alphabetic characters
func CleanQuery(query string) []byte {
	var res []byte
	for i := range query {
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
	// Define variables
	var (
		bestMatch      []byte
		bestMatchScore float64 = -1.0
	)

	// Iterate over the subject names
	for subjectName := range global.SubjectNames {
		var (
			// Convert subject name to bytes
			subjectName []byte = []byte(subjectName)

			// Result value
			score float64 = 0.0

			// Get the largest / smallest keys
			largestKey  []byte = GetLargest(subjectName, query)
			smallestKey []byte = GetSmallest(subjectName, query)
		)

		// Iterate using the smallest key length
		for i := range smallestKey {
			var (
				tempIndex float64 = 1.0
				substr    []byte
			)

			// If the keys equal the same
			if query[i] == subjectName[i] {
				score += float64(query[i])
			} else {
				score -= float64(query[i]) / float64(len(largestKey))
			}

			// Iterate over the smallest key
			for j := range smallestKey {
				// Add the letter to the substr
				substr = append(substr, query[j])

				// Check if the subjectName contains the substr
				if bytes.Contains(subjectName, substr) {
					// Make sure the length of the contain check
					// is greater than 2, or else you'll use single letters
					if len(substr) > 2 {
						score += float64(query[j]) / float64(len(substr))
					}
				} else {
					// Reset the contain check
					substr = []byte{}
				}

				// Get the distance the same letters are from eachother
				// using the tempIndex
				tempIndex++
				if subjectName[i] == smallestKey[j] {
					score += tempIndex / float64(len(smallestKey)*len(largestKey))
				}
			}
		}

		// Check if score is greater than
		// the previous bestMatchScores
		if score > bestMatchScore {
			bestMatchScore = score
			bestMatch = subjectName
		}
	}

	// Make sure best match is valid/accurate
	if bestMatchScore > float64(370-(len(bestMatch)/2)) {
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
func QueryHandler(query string) string {
	// Convert the query to lowercase
	query = strings.ToLower(query)

	// Check if the user is searching for a specific subject code
	var splitQuery []string = strings.Split(query, "@code")
	if len(splitQuery) > 1 {
		return string(CleanQuery(splitQuery[1]))
	}

	// If using a search query (ex: computerscience)
	// then match the query to a subject code
	var cleanedQuery []byte = []byte(CleanQuery(query))
	return string(GetBestMatch(cleanedQuery))
}
