package api

// Import modules
import (
	"sync"

	hermes "github.com/realTristan/Hermes/nocache"
)

// GetCourses() returns the courses from the cache
func GetCourses(ft *hermes.FullText, query string, subject string) []map[string]interface{} {
	var (
		wg            *sync.WaitGroup = &sync.WaitGroup{}
		queryResult   []map[string]interface{}
		subjectResult []map[string]interface{}
	)

	// Query the courses
	wg.Add(1)
	go func() {
		queryResult, _ = ft.Search(query, 100, false, map[string]bool{
			"id":             false,
			"components":     false,
			"units":          false,
			"description":    true,
			"name":           true,
			"pre_requisites": true,
			"title":          true,
		})
		wg.Done()
	}()

	// Search for the course title
	if len(queryResult) <= 100 {
		wg.Add(1)
		go func() {
			subjectResult, _ = ft.SearchValuesWithKey(subject, "title", 100)
			wg.Done()
		}()
	}
	wg.Wait()

	// Return the two merged queries
	return append(subjectResult, queryResult...)
}
