package redis

// Import modules
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

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
	var res []string = strings.Split(query, " ")
	res = append(res, strings.ToLower(query))
	return res
}

// The GetSimilarCourses() function iterates through the redis
// cache and gets any courses that contain the query args
func GetSimilarCourses(html string, query string) (string, int) {
	query = strings.TrimSpace(query)

	// Define Variables
	var (
		// queryArgs: []string -> Used to find similar courses
		queryArgs []string = GetQueryArgs(query)

		// resultAmount: int -> The amount of similar courses
		resultAmount int = 0

		// WwaitGroup: sync.WaitGroup -> wait group for goroutines
		waitGroup sync.WaitGroup = sync.WaitGroup{}
	)

	// Iterate over all the keys in the database
	for _, key := range GetAllKeys() {
		waitGroup.Add(1)

		// Goroutine for faster query time
		go func(key interface{}) {
			defer waitGroup.Done()

			// Json unmarshal the json encoded map
			var data []map[string]string
			json.Unmarshal([]byte(Get(key.(string))), &data)

			// For every query arg check if the
			// map contains the arg
			for q := 0; q < len(queryArgs); q++ {
				for v := 0; v < len(data); v++ {
					// Check if the data contains the queryArg
					if strings.Contains(
						strings.ToLower(fmt.Sprint(data[v])), " "+queryArgs[q]+" ") {

						// Check if course is already present
						if !strings.Contains(html, data[v]["ID"]) {
							resultAmount++

							// Append to the html string
							html += GenerateCourseHTML(data[v])
						}
					}
				}
			}
		}(key)
	}
	// Wait for all goroutines
	waitGroup.Wait()

	// Return the html, resultAmount
	return html, resultAmount
}

// The GenerateCourseHTML() function will use the course
// data map to generate an html string that is used for showing
// the list of courses on the home page.
func GenerateCourseHTML(data map[string]string) string {
	// Define Variables
	var (
		// Result: string -> The html result
		result string = ""

		// keys: []string -> They keys from the course data map
		keys []string = []string{
			"Components", "Unit", "ID", "Name",
			"Description", "Notes", "Pre Requisites",
			"Anti Requisites",
		}
	)

	// Iterate over the keys and add them
	// to the html result. The html result
	// will be returned inside styled divs
	for i := 0; i < len(keys); i++ {
		if len(data[keys[i]]) > 0 {
			result += fmt.Sprintf(`
			<div style="font-size:13px;">
				<strong>%s</strong> %v
			</div>
			`, keys[i], data[keys[i]])
		}
	}

	// Return the html wrapped in styled divs
	// If anyone has any suggestions for this part of the
	// code, please tell me!
	return fmt.Sprintf(`
		<div style="width: 100%%; margin-top: 5%%; margin-bottom: -1.3%%">
			<div class="course_div" style="position: relative;">
			<img 
				src="static/images/waterloo_logo.png" 
				alt="" 
				width="80"
				height="80"
				style="position: absolute; top: 5%%; right: 1%%;"
			>
				<div style="font-size:13px;">
					<strong>%v</strong>
				</div>
				%s
			</div>
		</div>`, data["Title"], result)
}
