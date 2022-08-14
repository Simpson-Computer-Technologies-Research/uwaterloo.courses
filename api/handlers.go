package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/realTristan/The_University_of_Waterloo/global"
	scraper "github.com/realTristan/The_University_of_Waterloo/scraper"
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
			course      string = QueryHandler(r)
			result, err        = scraper.ScrapeCourseData(RequestClient, strings.ToUpper(course))
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
		// Find out how to create a list with all the data
		// This is the page that handles the query search
		//
		// Use the university waterloo text logo as the header for the main page
		// Use the university of waterloo badge logo inside the course data list
		//
		// For the home page learn svelte with go and typescript
		// and tailwind css
		//
		// Set the content type
		w.Header().Add("Content-Type", "text/html")
		// Execute the html template
		Template.Execute(w, nil)
	}
}

// The DevTestingHandler() function is used for developement testing
func DevTestingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Developement Testing Endpoint")
	}
}
