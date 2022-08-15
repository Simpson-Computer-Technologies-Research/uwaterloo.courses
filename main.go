package main

// Import packages
import (
	"github.com/realTristan/The_University_of_Waterloo/api"
)

// Main function
func main() {
	// scraper.RefreshCache()
	api.ListenAndServe(":8000")
}

/*

# UNIVERSITY OF WATERLOO COURSE FINDER

WHAT TO DO NEXT:
- Learn Redis-Search module for querying
- Make Frontend look better (Svelte or React)
- Learn Docker


FOR SVELTE:
https://www.sitepoint.com/svelte-fetch-data/
- fetch data from localhost api (golang api)


PROJECT NOTES:

// DOCKER
// Learn docker and try to use svelte with it

// SVELTE / FRONTEND
// Learn svelte and try to use svelte for the frontend
// Use svelte to make frontend look better

// SEARCHING
// Use redis for searching which will make it much
// faster

// HOSTING
// host with fly.io

*/
