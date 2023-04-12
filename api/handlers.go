package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"net/http"
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

		// Get the query
		var query = r.URL.Query().Get("q")

		// Make sure the query length is greater than 3
		if len(query) < 3 {
			w.Write([]byte("[]"))
			return
		}

		// Define the variables
		var (
			startTime time.Time = time.Now()
			subject   string    = QueryHandler(query)
		)

		// Marhsal the response
		var resp, _ = json.Marshal(map[string]interface{}{
			"query":  query,
			"result": cache.GetCourses(query, subject),
			"time":   time.Since(startTime).Microseconds(),
		})

		// Set the response body
		w.Write(resp)
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
