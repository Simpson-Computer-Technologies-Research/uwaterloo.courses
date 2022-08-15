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
			course         string = QueryHandler(r)
			result, _, err        = scraper.ScrapeCourseData(RequestClient, strings.ToUpper(course))
		)

		// Handle the error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		} else {
			// Marshal the data result
			_json, _ := json.Marshal(map[string]interface{}{
				course: result,
			})

			// Set the response body
			fmt.Fprint(w, string(_json))
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
		w.Header().Add("Content-Type", "text/html")

		// If there's a query arg, return the scraped data
		// in html format
		//
		// Else, return the home page with the search bar
		if len(r.URL.Query().Get("q")) > 0 {
			var (
				course string = QueryHandler(r)
				result []map[string]string
				html   string
			)
			// If the course key is not in the redis database
			// then run the scraper to get the course data
			if !redis.Contains(course) {
				result, html, _ = scraper.ScrapeCourseData(RequestClient, strings.ToUpper(course))
			} else {
				// Else, if the course key is in the redis
				// database, then generate an html string
				// and set the html variable to said string
				json.Unmarshal([]byte(redis.Get(course)), &result)
				html = redis.GenerateHTML(result)
			}

			// Execute the scraped data page html template
			Template.Execute(w,
				fmt.Sprintf("%s%s", global.EndQueryTimer(len(result)), html))
		} else {
			// Execute the home page html template
			Template.Execute(w, nil)
		}
	}
}

// The DevTestingHandler() function is used for developement testing
func DevTestingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Developement Testing Endpoint")
	}
}
