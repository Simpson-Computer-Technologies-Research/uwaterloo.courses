package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/realTristan/The_University_of_Waterloo/global"
	"github.com/realTristan/The_University_of_Waterloo/redis"
	"github.com/realTristan/The_University_of_Waterloo/scraper"
	"github.com/valyala/fasthttp"
)

// Global Variables
/* - RequestClient *fasthttp.Client -> Used for sending http requests */
/* - Template *template.Template -> Used for rendering html templates */
var (
	RequestClient *fasthttp.Client   = &fasthttp.Client{}
	Template      *template.Template = template.Must(template.ParseGlob("static/templates/*.html"))
)

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
				// Set the response body
				fmt.Fprint(w, "Query needs to be greater than 3 characters!")
			} else {
				// Marshal the redis course json data
				json.Unmarshal([]byte(redis.Get(strings.ToUpper(subject))), &result)

				// Marshal the data result
				redis.GetSimilarCourses(&result, "", query)

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
//
// The function renders the index.html file
func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		global.StartQueryTimer()

		// Set the header for html
		w.Header().Add("Content-Type", "text/html")

		// Set the query: string variable
		var query string = strings.ToLower(r.URL.Query().Get("q"))

		// If there's a query arg, return the scraped data
		// in html format
		//
		// Else, return the home page with the search bar
		if len(query) > 0 {
			var (
				// course: string -> The course code
				course string = QueryHandler(r)

				// result: []map[string]string -> The course data result map
				result []map[string]string

				// html: string -> The result html
				html string

				// resultAmount: int -> The amount of similar courses to add
				resultAmount int = 0
			)
			// If the course key is not in the redis database
			// then run the scraper to get the course data
			if !redis.Exists(course) && len(course) > 0 {
				// Scrape the course data
				result, _ = scraper.ScrapeCourseData(RequestClient, strings.ToUpper(course))

				// Iterate over the result slice
				for i := 0; i < len(result); i++ {
					// Append the generated html to the result html
					html += global.GenerateCourseHTML(result[i])
				}
			} else {
				// Check to make sure the course isn't empty
				if len(course) > 0 {
					// Unmarshal the course data from the
					// redis cache
					json.Unmarshal([]byte(redis.Get(course)), &result)

					// Iterate over the result slice
					for i := 0; i < len(result); i++ {
						// Append to the result html, the generated html
						html += global.GenerateCourseHTML(result[i])
					}
				}
				// Create a temp result map for the get similar courses function
				var tempResultMap []map[string]string = []map[string]string{}

				// Get the similar courses using query keywords
				_, html, resultAmount = redis.GetSimilarCourses(
					&tempResultMap, html, query)
			}

			// Execute the scraped data page html template
			Template.Execute(w,
				fmt.Sprintf("%s%s",
					global.QueryMenu(len(result)+resultAmount), html))
		} else {
			// Execute the home page html template
			Template.Execute(w, global.HomePageSearchBar())
		}
	}
}
