package cache

// Import modules
import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

/*

	Some of you might be wondering why I decided to use a byte cache
	instead of a map cache. For starters iterating over a map cache takes
	too long. To solve that you can start a bunch of goroutines which outputs
	the same speed but destroys your memory usage. (especially if a lot of
	users are calling the api) Secondly, in most cases the byte cache is
	faster.

*/

// Hold Course data in memory cache
var Cache []byte

// The Set() function sets the data for the
// given key in the cache
func Set(value map[string]string) {
	var tmp, _ = json.Marshal(value)

	// Append to a byte array is faster than
	// adding to a string. This makes the webscraping
	// part of the program much faster (30x faster)
	Cache = append(Cache, tmp...)
}

// The GetCourses() function iterates through the
// cache and gets any courses that contain the query
// as well as any courses that start with the subject code
func GetCourses(query []byte, subject []byte) []map[string]string {
	subject = []byte(fmt.Sprintf(`,"title":"%s `, subject))

	// Define variables
	var (
		// Track query time
		startTime time.Time = time.Now()

		// courseMapStart -> Track opening bracket
		courseMapStart int = -1

		// closeBracketCount -> Track closing brackets per course map
		closeBracketCount int = 0

		// subjectResult -> Array with all the courses that have the subject code
		subjectResult []map[string]string

		// similarResult -> Array with all courses that contain the query
		similarResult []map[string]string

		// TempCache -> Lowercase Cache string
		TempCache []byte = bytes.ToLower(Cache)
	)

	// Iterate over the lowercase cache string
	for i := 0; i < len(TempCache); i++ {
		// Break the loop if there's too many similar courses
		if len(similarResult) > 500 {
			break
		} else

		// Check if current index is the start of
		// the course data map
		if TempCache[i] == '{' {
			if courseMapStart == -1 {
				courseMapStart = i
			}
			closeBracketCount++
		} else

		// Check if the current index is the end of
		// the course data map
		if TempCache[i] == '}' {
			if closeBracketCount == 1 {
				// Check if the map contains the subject code
				if bytes.Contains(Cache[courseMapStart:i+1], subject) {

					// Convert the string to a map
					var data map[string]string
					json.Unmarshal(Cache[courseMapStart:i+1], &data)

					// Append the map to the result array
					subjectResult = append(subjectResult, data)
				} else

				// Check if the map contains the query string
				if bytes.Contains(TempCache[courseMapStart:i+1], query) {
					// Convert the string to a map
					var data map[string]string
					json.Unmarshal(Cache[courseMapStart:i+1], &data)

					// Append the map to the result array
					similarResult = append(similarResult, data)
				}
				// Reset indexing variables
				closeBracketCount = 0
				courseMapStart = -1
			} else {
				closeBracketCount--
			}
		}
	}
	// Print the query time
	fmt.Printf("\n >> Course Query: (%v)\n", time.Since(startTime))

	// Return the combined arrays
	return append(subjectResult, similarResult...)
}
