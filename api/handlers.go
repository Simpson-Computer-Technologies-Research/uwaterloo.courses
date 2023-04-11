package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/realTristan/uwaterloo.courses/cache"
	"github.com/realTristan/uwaterloo.courses/global"
)

// The CourseDataHandler() function handles the incoming requests
// using the "/courses?q={query}" path.
// The function is used to scrape the data of a subject using the
// ScrapeCourseData() function, then return it as a json string
func CourseDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Define variables
		var (
			startTime      time.Time = time.Now()
			query, subject           = QueryHandler(r)
		)

		// Make sure the query length is greater than 3
		if len(query) >= 3 {
			// Convert the subject to lowercase
			subject = strings.ToLower(subject)

			// Generate the response
			var resp, _ = json.Marshal(map[string]interface{}{
				"query_time": time.Since(startTime).Microseconds(),
				"courses":    cache.GetCourses(query, subject),
			})
			w.Write(resp)
		} else {
			w.Write([]byte("[]"))
		}
	}
}

// The SubjectCodesHandler() function handles the incoming
// requests using the "/subjects" path
//
// The function returns the global SubjectCodes array
func SubjectCodesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Marshal the subject codes map
		_json, _ := json.Marshal(map[string][]string{
			"subjects": global.SubjectCodes,
		})

		// Set the response body
		w.Write(_json)
	}
}

// The SubjectCodesWithNamesHandler() function handles the incoming
// requests using the "/subjects/names" path
//
// The function returns the global SubjectNames
func SubjectCodesWithNamesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Marshal the codes and names map
		_json, _ := json.Marshal(global.SubjectNames)

		// Set the response body
		w.Write(_json)
	}
}

// The HomePageHandler() function handles the incoming requests
// using the "/" path.
func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "uwaterloo.courses Public API")
	}
}

// Dev page handler
func DevPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
