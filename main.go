package main

// Import packages
import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// Main function
func main() {
	// Function start time
	var startTime = time.Now()

	// Define Variables
	var (
		// RequestClient to use for sending htp requests
		RequestClient *fasthttp.Client = &fasthttp.Client{}
		// The Scraped course codes
		codes, err = ScrapeSubjectCodes(RequestClient)
	)
	// Print the result
	fmt.Printf("%v: %v: %v", time.Since(startTime), err, codes)
}

/*

# UNIVERSITY OF WATERLOO COURSE FINDER

// Get the course title with the Yellow Highlighted Title at the top of the page

>> For holding data in https://ucalendar.uwaterloo.ca/2021/COURSE/course-CS.html
>> Map Type: map[string][interface{}]

// Map Data
"Computer Science": {
	"course id": {
		"info": "CS 100 LAB,LEC,TST,TUT 0.50",
		"course_id": 012765,
		"title": "Introduction to Computing through Applications",
		"desc": "Using personal computers as effective problem solving tools for...",
		"pre_req": "Prereq: Not open to Mathematics,Biomedical...",
		"anti_req": "All second,third or fourth year CS courses or equivalents"
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
