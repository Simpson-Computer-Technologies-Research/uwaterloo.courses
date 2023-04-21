package api

// Import modules
import hermes "github.com/realTristan/Hermes"

// GetCourses() returns the courses from the cache
func GetCourses(ft *hermes.FullText, query string, subject string) []map[string]string {
	var (
		// Search for the course title
		subjectResult []map[string]string = ft.SearchInJsonWithKey(subject, "title", 100)

		// Search for query variable
		queryResult []map[string]string = ft.SearchWithSpaces(query, 100, false, map[string]bool{
			"id":             false,
			"components":     false,
			"units":          false,
			"description":    true,
			"name":           true,
			"pre_requisites": true,
			"title":          true,
		})
	)

	// Return the two merged queries
	return append(subjectResult, queryResult...)
}
