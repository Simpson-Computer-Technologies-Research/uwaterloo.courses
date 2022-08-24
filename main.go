package main

import (
	"github.com/realTristan/The_University_of_Waterloo/server/api"
	"github.com/realTristan/The_University_of_Waterloo/server/scraper"
)

// Main function
func main() {
	// Refresh the course info database
	scraper.RefreshCache()

	// Host the API
	api.ListenAndServe(":8000")
}

/*

# UNIVERSITY OF WATERLOO COURSE FINDER

WHAT TO DO NEXT:
- Learn Docker
- Learn Redis Search for faster querying
- Host the website using fly.io

*/
