package api

import (
	"strings"
	"unicode"

	"github.com/realTristan/uwaterloo.courses/global"
)

// The CleanQuery() function removes all spaces from the query
// and removes all non alphabetic characters
func CleanQuery(query string) string {
	var res []byte = []byte{}
	for i := range query {
		// Check if the character at the index is a letter
		if unicode.IsLetter(rune(query[i])) {
			// Append the letter to the result
			res = append(res, query[i])
		}
	}
	return string(res)
}

// The GetLargest() function returns the
// largest byte out of the two
func GetLargest(a string, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

// The GetSmallest() function returns the
// smallest byte out of the two
func GetSmallest(a string, b string) string {
	if len(a) > len(b) {
		return b
	}
	return a
}

// The GetBestMatch() function uses the cleaned query (ex: computerscience)
// to find the best match using the global.Subjects map
// It returns the best subject code match (ex: CS)
func GetBestMatch(query string) string {
	// If the query is in the subject names
	if _, ok := global.Subjects[query]; ok {
		return global.Subjects[query]
	}

	// Define variables
	var (
		bestMatch      string
		bestMatchScore float64 = -1.0
	)

	// Iterate over the subject names
	for _, subjectName := range global.SubjectNames {
		switch {
		// If the subjectName length is less than the query name length
		case len(subjectName) < len(query):
			continue

		// If the subjectName starts with the prefix
		case strings.HasPrefix(subjectName, query):
			return global.Subjects[subjectName]
		}

		// Define variables
		var (
			score       float64 = 0.0
			largestKey  string  = GetLargest(subjectName, query)
			smallestKey string  = GetSmallest(subjectName, query)
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

				// If the subjectName contains the substr
				if strings.Contains(subjectName, string(substr)) {
					// Add to the score
					if len(substr) > 2 {
						score += float64(query[j]) / float64(len(substr))
					}
				} else {
					// Reset the substring
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
		return global.Subjects[bestMatch]
	}
	return ""
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
		return CleanQuery(splitQuery[1])
	}

	// If using a search query (ex: computerscience)
	// then match the query to a subject code
	return GetBestMatch(CleanQuery(query))
}
