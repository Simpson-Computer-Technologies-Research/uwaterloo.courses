package redis

// Import modules
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v9"
	"github.com/realTristan/The_University_of_Waterloo/server/global"
)

// RUN docker run --name redis-test-instance -p 6379:6379 -d redis

// Define Variables
var (
	// Context: context.Context -> Used for the Redis Cache
	Context context.Context = context.Background()

	// RedisCache: *redis.Client -> Used to Connect to the redis database
	RedisCache *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

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
// or not the redis cache contains the given key
func Exists(key string) bool {
	_, err := RedisCache.Get(Context, key).Result()
	return err != redis.Nil
}

// The Set() function sets the data for the
// given key in the redis cache
func Set(key string, value string) error {
	return RedisCache.Set(Context, key, value, 0).Err()
}

// The Get() function returns the data
// for the given key from the redis cache
func Get(key string) string {
	var v, _ = RedisCache.Get(Context, key).Result()
	return v
}

// The GetSimilarCourses() function iterates through the redis
// cache and gets any courses that contain the query args
func GetSimilarCourses(rsc *SimilarCourses) []map[string]string {
	rsc.Query = strings.ToLower(strings.TrimSpace(rsc.Query))

	// WaitGroup: sync.WaitGroup -> wait group for goroutines
	var waitGroup sync.WaitGroup = sync.WaitGroup{}

	// Iterate over all the keys in the database
	for i := 0; i < len(global.SubjectCodes); i++ {
		waitGroup.Add(1)

		go func(key interface{}) {
			defer waitGroup.Done()

			// Make sure not to add duplicate courses
			if key != rsc.Subject {

				// Json unmarshal the json encoded map
				var data []map[string]string
				json.Unmarshal([]byte(Get(key.(string))), &data)

				// For every query arg check if the
				// map contains the arg
				if strings.Contains(fmt.Sprint(data), rsc.Query) {
					for v := 0; v < len(data); v++ {
						go func(v int) {
							// Check if the data contains the queryArg
							if strings.Contains(strings.ToLower(fmt.Sprint(data[v])), rsc.Query) {
								// Append data to the result array
								rsc.ResultArray = append(rsc.ResultArray, data[v])
							}
						}(v)

						// Break the loop if the result array is
						// too large
						if len(rsc.ResultArray) > 500 {
							break
						}
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
