package api

// Import modules
import (
	hermes "github.com/realTristan/Hermes"
)

// Hermes cache
var cache *hermes.Cache = hermes.InitCache("default_data.json")

// GetCourses() returns the courses from the cache
func GetCourses(query string, subject string) []map[string]string {
	var (
		// Search for the course title
		subjectResult, _ = cache.SearchInJsonWithKey(subject, "title", 100, false)

		// Search for query variable
		queryResult, _ = cache.SearchWithSpaces(query, 100, false, []string{
			"id",
			"units",
			"components",
		})
	)

	// Return the two merged queries
	return append(queryResult, subjectResult...)
}
