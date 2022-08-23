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

// The SimilarCourses struct holds three keys
/* - ResultArray: []map[string]string -> result data		*/
/* - Subject: string -> the handled subject					*/
/* - Query: string -> the original query					*/
type SimilarCourses struct {
	ResultArray []map[string]string
	Subject     string
	Query       string
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

// The GetAllKeys() function returns all the
// keys from the redis cache
func GetAllKeys() []interface{} {
	keys, _ := RedisCache.Do(Context, "KEYS", "*").Result()
	return keys.([]interface{})
}

// The GetSimilarCourses() function iterates through the redis
// cache and gets any courses that contain the query args
func GetSimilarCourses(rsc *SimilarCourses) []map[string]string {
	rsc.Query = strings.ToLower(strings.TrimSpace(rsc.Query))

	// WaitGroup: sync.WaitGroup -> wait group for goroutines
	var waitGroup sync.WaitGroup = sync.WaitGroup{}

	// Iterate over all the keys in the database
	for _, key := range GetAllKeys() {
		go func(key interface{}) {

			// Json unmarshal the json encoded map
			var data []map[string]string
			json.Unmarshal([]byte(strings.ToLower(Get(key.(string)))), &data)

			// Get the amount of iterations to perform
			var (
				rec     int = len(data)
				pre_rec int = rec * (len(rsc.Query) / rec)
			)
			// Set the amount of iterations to pre_rec
			if pre_rec < rec {
				rec = pre_rec
			}

			// For every query arg check if the
			// map contains the arg
			for v := 0; v < rec; v++ {
				waitGroup.Add(1)

				go func(v int) {
					defer waitGroup.Done()

					// Check if the data contains the queryArg
					if strings.Contains(fmt.Sprint(data[v]), rsc.Query) {

						// Check if course is already present
						if !strings.Contains(fmt.Sprint(rsc.ResultArray), data[v]["ID"]) {
							rsc.ResultArray = append(rsc.ResultArray, data[v])
						}
					}
				}(v)
			}
		}(key)
	}
	// Wait for all goroutines
	waitGroup.Wait()

	// Return the result array, resultAmount
	return rsc.ResultArray
}
