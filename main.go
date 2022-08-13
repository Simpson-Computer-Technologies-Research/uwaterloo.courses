package main

// Import packages
import (
	"encoding/json"
	"fmt"

	"github.com/realTristan/The_University_of_Waterloo/global"
	scraper "github.com/realTristan/The_University_of_Waterloo/scraper"
	"github.com/valyala/fasthttp"
)

// Fasthttp request client
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// Function to handle incoming http requests
func CourseDataHandler(ctx *fasthttp.RequestCtx) {
	// Scrape the course data
	var (
		course      []byte = ctx.QueryArgs().Peek("course")
		result, err        = scraper.ScrapeCourseData(RequestClient, string(course))
	)
	// Handle the error
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprintf(ctx, "{\"error\": \"%v\"}", err)
	} else {
		// Marshal the data result
		_json, _ := json.Marshal(result)

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

// Main function
func main() {
	// Localhost port
	var port string = ":8080"

	// Print the localhost url
	fmt.Printf("Listening on: http://localhost%s", port)

	// Listen and Server the port
	fasthttp.ListenAndServe(port, func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		// Show course data with the paramter ?course={course_code}
		case "/courses":
			CourseDataHandler(ctx)

		// Show the list of subjects at the university of waterloo
		case "/subjects":
			SubjectCodesHandler(ctx)

		// Show the list of subjects with their corresponding names
		// at the university of waterloo
		case "/subjects/names":

		// Invalid path error
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	})
}

/*
# UNIVERSITY OF WATERLOO COURSE FINDER

WHAT TO DO NEXT:

// Learning
- Learn Redis
- Learn Svelte

// Doing
- Cache data into redis db
- Create Frontend


PROJECT NOTES:

// Get the course title with the Yellow Highlighted Title at the top of the page
>> Course Data: https://ucalendar.uwaterloo.ca/2021/COURSE/course-CS.html

// Map Data
"Computer Science": {
	"course_id": {
		"title": "CS 100",
		"components": LAB,LEC,TST,TUT",
		"unit": "0.50",
		"id": 012765,
		"title": "Introduction to Computing through Applications",
		"desc": "Using personal computers as effective problem solving tools for...",
		"pre_req": "Prereq: Not open to Mathematics,Biomedical...",
		"anti_req": "All second,third or fourth year CS courses or equivalents",
		"other": "online url"
	}
}

// PRE REQUISITE MAPPING
// Let's say another course requires CS 100 as a pre_req, it'll list each
// course that requires CS 100 as a pre_req

// MEMORY USAGE
// if the memory usage is too high which it probably will be
// use redis a database system and have a temp map caching system

// MAP CACHING SYSTEM
// have a maximum sized map for example only 100 keys
// and it'll remove the key from the bottom
// and add the new map key to the front

// FRONTEND
// finish the backend and the mapping, caching and database systems
// before creating a frontend
//
// Style Ideas:
// background is the red color of the uni of waterloo logo

// DATA API
// Create an api that will return data from the redis database
// for anyone to use

// HOSTING
// host with fly.io
*/
