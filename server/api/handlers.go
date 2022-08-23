package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/realTristan/The_University_of_Waterloo/server/global"
	"github.com/realTristan/The_University_of_Waterloo/server/redis"
	"github.com/valyala/fasthttp"
)

// Global Variables
/* - RequestClient *fasthttp.Client -> Used for sending http requests */
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// The CourseDataHandler() function handles the incoming requests
// using the "/courses?course={course_code}" path.
// The function is used to scrape the data of a subject using the
// ScrapeCourseData() function, then return it as a json string
func CourseDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Define Variables
		// course: string -> Get the course to search for
		// result, err -> Scrape the course data
		var (
			subject string = QueryHandler(r)
			err     error  = nil
			result  []map[string]string
		)

		// Handle the error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		} else {
			// Enable CORS
			w.Header().Set("Access-Control-Allow-Origin", "*")

			// Set the query: string variable
			var query string = strings.ToLower(r.URL.Query().Get("q"))

			// Make sure the query length is greater than 3
			if len(query) < 3 {
				// Return the query error
				fmt.Fprint(w, "Query needs to be greater than 3 characters!")
			} else {
				// Marshal the redis course json data
				json.Unmarshal([]byte(redis.Get(strings.ToUpper(subject))), &result)

				// Get the similar courses from the redis database
				var rsc *redis.SimilarCourses = &redis.SimilarCourses{
					ResultArray: result,
					Subject:     subject,
					Query:       query,
					Mutex:       &sync.RWMutex{},
				}
				result = redis.GetSimilarCourses(rsc)

				// Convert the result map into a json string
				var res, _ = json.Marshal(result)

				// Set the response body
				fmt.Fprint(w, string(res))
			}
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
		fmt.Fprint(w, string(_json))
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
		fmt.Fprint(w, string(_json))
	}
}

// The HomePageHandler() function handles the incoming requests
// using the "/" path.
func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "uwaterloo.courses Public API")
	}
}
