package main

import (
	"github.com/realTristan/The_University_of_Waterloo/server/api"
)

// Main function
func main() {
	// Refresh the course info database
	// scraper.RefreshCache()

	// From redis database to in memory cache
	// cache.FromRedisToCache()

	// Host the API
	api.ListenAndServe(":8000")

	// Speed Testing
	// cache.TestGetSimilarCourses()
}

/*

# UNIVERSITY OF WATERLOO COURSE FINDER

WHAT TO DO NEXT:
- Learn Docker
- Learn Redis Search for faster querying
- Host the website using fly.io

*/
