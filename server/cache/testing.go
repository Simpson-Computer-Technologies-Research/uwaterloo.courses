package cache

// Import packages
import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/realTristan/The_University_of_Waterloo/server/global"
	"github.com/realTristan/The_University_of_Waterloo/server/redis"
)

// The FromRedisToCache() function takes all the redis data
// and adds it to the in memory cache
func FromRedisToCache() {
	for _, key := range global.SubjectCodes {
		// Unmarshal the data
		data := []map[string]string{}
		json.Unmarshal([]byte(redis.Get(key)), &data)

		// Set course data
		Cache[key] = data
	}
}

// The TestSimilarCourses() function is used to test
// the GetSimilarCourses() function speed
//
// The function gets data from the redis database
// then adds it all to the in memory cache
// from ./cache.go
func TestGetSimilarCourses() {
	// Start time for testing speed
	startTime := time.Now()

	// Create Similar Courses object
	var rsc *SimilarCourses = &SimilarCourses{
		ResultArray: []map[string]string{},
		Subject:     "CS",
		Query:       "science",
		Mutex:       &sync.RWMutex{},
	}
	// Get similar courses
	GetSimilarCourses(rsc)

	// Print the amount of similar courses and speed
	fmt.Printf("Result: (%d) (%v)\n", len(rsc.ResultArray), time.Since(startTime))
}
