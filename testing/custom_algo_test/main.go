package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Initialize the cache
var cache map[string][]int = make(map[string][]int)
var cacheKeys []string = []string{}

// Initialize the json data
var jsonData []map[string]string = []map[string]string{}

// Main function
func main() {
	// Load the cache
	LoadCache()

	// Listen and serve on port 8000
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8000", nil)
}

// Handle the incoming http request
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the query parameter
	var query string = r.URL.Query().Get("q")

	// Track the start time
	var start time.Time = time.Now()

	// Search for a word in the cache
	var courses []map[string]string = IndicesToCourses(Search(query))

	// Print the duration
	fmt.Println(time.Since(start))

	// Write the courses to the json response
	var response, _ = json.Marshal(courses)
	w.Write(response)
}

// Search for a word in the cache
func Search(word string) []int {
	word = strings.ToLower(word)
	var result []int = []int{}
	for i := 0; i < len(cacheKeys); i++ {
		if strings.Contains(cacheKeys[i], word) {
			result = append(result, cache[cacheKeys[i]]...)
		}
	}
	return result
}

// Indices to courses
func IndicesToCourses(indices []int) []map[string]string {
	var courses []map[string]string = make([]map[string]string, 0)
	for i := 0; i < len(indices); i++ {
		courses = append(courses, jsonData[indices[i]])
	}
	return courses
}

// Load the cache
func LoadCache() {
	// Read the json file
	var data, _ = os.ReadFile("default_data.json")
	json.Unmarshal(data, &jsonData)

	// Loop through the json data
	for i, item := range jsonData {
		for _, runeValue := range item {
			// Convert the rune to a string
			var value string = string(runeValue)

			// Remove double spaces
			for strings.Contains(value, "  ") {
				value = strings.Replace(value, "  ", " ", -1)
			}

			// Remove spaces from front and back of the string and convert to lowercase
			value = strings.ToLower(strings.TrimSpace(value))

			// Split the string into an array
			var array []string = strings.Split(value, " ")

			// Loop through the array
			for _, word := range array {
				if _, ok := cache[word]; !ok {
					cache[word] = []int{}
				}

				// Update the cache keys
				if !Contains(cacheKeys, word) {
					cacheKeys = append(cacheKeys, word)
				}
				cache[word] = append(cache[word], i)
			}
		}
	}
}

// Check if a string is in an array
func Contains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}