package redis

// Import modules
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v9"
	"github.com/realTristan/The_University_of_Waterloo/global"
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

// The GetAllKeys() function returns all the
// keys from the redis cache
func GetAllKeys() []interface{} {
	keys, _ := RedisCache.Do(Context, "KEYS", "*").Result()
	return keys.([]interface{})
}

// Get the query args from a query
// Example: query = computer science
// Returns [computer, science]
func GetQueryArgs(query string) []string {
	var res []string = strings.Split(query, " ")
	res = append(res, strings.ToLower(query))
	return res
}

// The GetSimilarCourses() function iterates through the redis
// cache and gets any courses that contain the query args
func GetSimilarCourses(result *[]map[string]string, html string, query string) (*[]map[string]string, string, int) {
	query = strings.ToLower(strings.TrimSpace(query))

	// Define Variables
	var (
		// resultAmount: int -> The amount of similar courses
		resultAmount int = 0

		// WwaitGroup: sync.WaitGroup -> wait group for goroutines
		waitGroup sync.WaitGroup = sync.WaitGroup{}
	)

	// Iterate over all the keys in the database
	for keyIndex, key := range GetAllKeys() {
		waitGroup.Add(1)

		// Goroutine for faster query time
		go func(key interface{}, keyIndex int) {
			defer waitGroup.Done()

			// Json unmarshal the json encoded map
			var data []map[string]string
			json.Unmarshal([]byte(Get(key.(string))), &data)

			// For every query arg check if the
			// map contains the arg
			var rec int = len(data)
			if len(data)*(len(query)/len(data)) < len(data) {
				rec = len(data) * (len(query) / len(data))
			}
			for v := 0; v < rec; v++ {
				waitGroup.Add(1)
				go func(v int) {
					defer waitGroup.Done()

					// Check if the data contains the queryArg
					if strings.Contains(
						strings.ToLower(fmt.Sprint(data[v])), query) {

						// Check if course is already present
						if !strings.Contains(html, data[v]["ID"]) {
							resultAmount++

							// Append to the html string
							html += global.GenerateCourseHTML(data[v])

							// Add data to the result map
							*result = append(*result, data[v])
						}
					}
				}(v)
			}
		}(key, keyIndex)
	}
	// Wait for all goroutines
	waitGroup.Wait()

	// Return the html, resultAmount
	return result, html, resultAmount
}
