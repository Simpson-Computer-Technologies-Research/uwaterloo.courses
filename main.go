package main

// Import packages
import (
	api "github.com/realTristan/The_University_of_Waterloo/api"
)

// Main function
func main() {
	api.ListenAndServe(":8000")
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
