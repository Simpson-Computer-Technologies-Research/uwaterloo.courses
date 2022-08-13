package api

// Import packages
import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/realTristan/The_University_of_Waterloo/global"
	scraper "github.com/realTristan/The_University_of_Waterloo/scraper"
	"github.com/valyala/fasthttp"
)

// Define Global Variables
var (
	RequestClient *fasthttp.Client   = &fasthttp.Client{}
	Template      *template.Template = template.Must(template.ParseGlob("public/templates/*.html"))
)

// The CourseDataHandler() function handles the incoming requests
// with the /courses?course={course_code} path.
// The function is used to scrape the data of a subject using the
// ScrapeCourseData() function, then return it as a json string
//
// The function takes the ctx *fasthttp.RequestCtx parameter
func CourseDataHandler(ctx *fasthttp.RequestCtx) {
	// Define Variables
	// course: string -> the course code arg
	// query: []byte -> the course search query arg
	var (
		course string = string(ctx.QueryArgs().Peek("course"))
		query  []byte = ctx.QueryArgs().Peek("q")
	)
	// If using a search query (ex: computerscience) then match the query
	// to a subject code
	if len(query) > 0 {
		course = SearchQuery(string(query))
	}

	// Scrape the course data
	var result, err = scraper.ScrapeCourseData(RequestClient, strings.ToUpper(course))

	// Handle the error
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprintf(ctx, "{\"error\": \"%v\"}", err)
	} else {
		// Marshal the data result
		_json, _ := json.Marshal(map[string]interface{}{
			course: result,
		})

		// Set the response body
		fmt.Fprint(ctx, string(_json))
	}
}

// The SubjectCodesHandler() function handles the incoming
// requests with the /subjects path
//
// The function returns the global SubjectCodes array
func SubjectCodesHandler(ctx *fasthttp.RequestCtx) {
	// Marshal the subject codes map
	_json, _ := json.Marshal(map[string][]string{
		"subjects": global.SubjectCodes,
	})

	// Set the response body
	fmt.Fprint(ctx, string(_json))
}

// The SubjectCodesWithNamesHandler() function handles the incoming
// requests with the /subjects/names path
//
// The function returns the global SubjectNames
func SubjectCodesWithNamesHandler(ctx *fasthttp.RequestCtx) {
	// Marshal the codes and names map
	_json, _ := json.Marshal(global.SubjectNames)

	// Set the response body
	fmt.Fprint(ctx, string(_json))
}

// The HomePageHandler() function handles the incoming requests
// using with / path.
//
// The function renders the index.html file
func HomePageHandler(ctx *fasthttp.RequestCtx) {
	// Set the content type
	ctx.Response.Header.Set("Content-Type", "text/html")
	// Execute the html template
	Template.Execute(ctx, nil)
}

// The DevTestingHandler() function is used for developement testing
func DevTestingHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Developement Testing Endpoint")
}
