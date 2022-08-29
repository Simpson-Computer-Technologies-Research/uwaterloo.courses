package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/realTristan/uwaterloo.courses/server/api"
	"github.com/realTristan/uwaterloo.courses/server/cache"
)

// The InitCache() function is used to read the
// default_data.json file and add all the course
// data from it to the cache.Cache array
func InitCache() {
	// Define Json File Variables
	var (
		jsonData, _  = os.Open("./default_data.json")
		jsonBytes, _ = io.ReadAll(jsonData)
	)

	// Unmarshal the JSON File Data
	var data []map[string]string
	json.Unmarshal(jsonBytes, &data)

	// Iterate through the json array
	for i := 0; i < len(data); i++ {

		// Marshal the json map
		var byteData, _ = json.Marshal(data[i])

		// Append the map to the cache.Cache array
		cache.Cache = append(cache.Cache, byteData)
	}
}

// Main function
func main() {
	// Refresh the course info database
	// scraper.RefreshCache()

	// Initialize the Cache
	InitCache()

	// Host the API
	api.ListenAndServe("0.0.0.0:8080")
}

/*

# UNIVERSITY OF WATERLOO COURSE FINDER

WHAT TO DO NEXT:
- Learn Docker
- Learn Redis Search for faster querying
- Host the website using fly.io

*/
