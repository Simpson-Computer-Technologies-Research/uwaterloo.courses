package cache

// Import modules
import (
	"fmt"
	"strings"
	"sync"

	"github.com/realTristan/The_University_of_Waterloo/server/global"
)

// Hold Course data in memory cache map
var Cache map[string][]map[string]string = make(map[string][]map[string]string)

// The SimilarCourses struct holds three keys
/* - ResultArray: []map[string]string -> result data		*/
/* - Subject: string -> the handled subject					*/
/* - Query: string -> the original query					*/
type SimilarCourses struct {
	ResultArray []map[string]string
	Subject     string
	Query       string
	Mutex       *sync.RWMutex
}

// The Exists() function checks whether
// or not the cache contains the given key
func Exists(key string) bool {
	var _, i = Cache[key]
	return i
}

// The Set() function sets the data for the
// given key in the cache
func Set(key string, value []map[string]string) error {
	Cache[key] = value

	// Add the data to the redis cache
	/*
		var data, _ = json.Marshal(value)
		return redis.Set(key, string(data))
	*/
	return nil
}

// The Get() function returns the data
// for the given key from the cache
func Get(key string) []map[string]string {
	return Cache[key]
}

// The GetSimilarCourses() function iterates through the
// cache and gets any courses that contain the query args
func GetSimilarCourses(rsc *SimilarCourses) []map[string]string {
	rsc.Query = strings.ToLower(rsc.Query)

	// WaitGroup: sync.WaitGroup -> wait group for goroutines
	var waitGroup sync.WaitGroup = sync.WaitGroup{}

	// Iterate over all the keys in the database
	for i := 0; i < len(global.SubjectCodes); i++ {
		waitGroup.Add(1)

		// Subject goroutine
		go func(key string) {
			defer waitGroup.Done()

			// Make sure not to add duplicate courses
			if key == rsc.Subject {
				return
			}
			// For every query arg check if the
			// map contains the arg
			if strings.Contains(fmt.Sprint(Cache[key]), rsc.Query) {
				for v := 0; v < len(Cache[key]); v++ {
					go func(v int) {
						// Check if the data contains the queryArg
						if strings.Contains(strings.ToLower(fmt.Sprint(Cache[key][v])), rsc.Query) {
							// Locking/Unlocking the mutex to prevent
							// data overwriting
							rsc.Mutex.Lock()
							defer rsc.Mutex.Unlock()

							// Append data to the result array
							rsc.ResultArray = append(rsc.ResultArray, Cache[key][v])
						}
					}(v)

					// Break the loop if the result array is
					// too large
					if len(rsc.ResultArray) > 500 {
						return
					}
				}
			}
		}(global.SubjectCodes[i])
	}
	// Wait for all goroutines
	waitGroup.Wait()

	// Return the result array, resultAmount
	return rsc.ResultArray
}
