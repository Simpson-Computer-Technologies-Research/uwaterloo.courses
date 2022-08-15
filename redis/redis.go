package redis

// Import modules
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
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
	return strings.Split(query, " ")
}

// The GetSimilarCourses() function...
func GetSimilarCourses(html string, query string) (string, int) {
	query = strings.TrimSpace(query)

	// Define Variables
	var (
		// queryArgs: []string -> Used to find similar courses
		queryArgs []string = GetQueryArgs(query)

		// resultAmount: int -> The amount of similar courses
		resultAmount int = 0
	)

	// Iterate over all the keys in the database
	for _, key := range GetAllKeys() {

		// Json unmarshal the json encoded map
		var data []map[string]string
		json.Unmarshal([]byte(Get(key.(string))), &data)

		// For every query arg check if the
		// map contains the arg
		for q := 0; q < len(queryArgs); q++ {
			// Iterate over the decoded map
			for v := 0; v < len(data); v++ {

				// Check if the data contains the queryArg
				if strings.Contains(fmt.Sprint(data[v]), queryArgs[q]) {

					// Check if course is already present
					if !strings.Contains(html, data[v]["ID"]) {
						// Add to the resultAmount
						resultAmount++

						// Append to the html string
						html += GenerateCourseHTML(data[v])

						/* Append the course title
						var title string = strings.Split(data[v]["title"], " ")[0]
						if !global.SliceContains(&queryArgs, title) {
							queryArgs = append(queryArgs, title)
						}*/
					}
				}
			}
		}
	}
	// Return the html, resultAmount
	return html, resultAmount
}
