package cache

// Import modules
import (
	"strings"

	Hermes "github.com/realTristan/Hermes"
)

// Initialize the hermes cache
var hermesCache *Hermes.Cache = Hermes.InitCache("default_data.json")

// Get the courses from the cache
func GetCourses(query string, subject string) []map[string]string {
	// Split the query into words
	var queryWords []string = strings.Split(query, " ")

	// Append the subject to the queryWords
	queryWords = append(queryWords, subject)

	// Search for the query in the cache
	return hermesCache.SearchMultiple(queryWords, 500, false)
}
