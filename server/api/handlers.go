package api

// Import packages
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/realTristan/uwaterloo.courses/server/cache"
	"github.com/realTristan/uwaterloo.courses/server/global"
	"github.com/valyala/fasthttp"
)

// Global Variables
/* - RequestClient *fasthttp.Client -> Used for sending http requests */
var RequestClient *fasthttp.Client = &fasthttp.Client{
	TLSConfig: &tls.Config{InsecureSkipVerify: true},
}

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
			subject []byte = QueryHandler(r)
			query   []byte = bytes.ToLower([]byte(r.URL.Query().Get("q")))
		)

		// Make sure the query length is greater than 3
		if len(query) >= 3 {
			// Get the courses
			var _res = cache.GetCourses(query, subject)
			res, _ := json.Marshal(_res)

			// Write the response
			w.Write(res)
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

// Dev page handler
func DevPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
