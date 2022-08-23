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
func GetSimilarCourses(rsc *SimilarCourses) []map[string]string {
	rsc.Query = strings.ToLower(strings.TrimSpace(rsc.Query))

	// WaitGroup: sync.WaitGroup -> wait group for goroutines
	var waitGroup sync.WaitGroup = sync.WaitGroup{}

	// Iterate over all the keys in the database
	for _, key := range GetAllKeys() {
		waitGroup.Add(1)

		// Goroutine for faster query time
		go func(key interface{}) {
			defer waitGroup.Done()

			// Json unmarshal the json encoded map
			var data []map[string]string
			json.Unmarshal([]byte(Get(key.(string))), &data)

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
				// Check if the data contains the queryArg
				if strings.Contains(
					strings.ToLower(fmt.Sprint(data[v])), rsc.Query) {

					// Check if course is already present
					if !strings.Contains(fmt.Sprint(rsc.ResultArray), data[v]["ID"]) {
						rsc.ResultArray = append(rsc.ResultArray, data[v])
					}
				}
			}
		}(key)
	}
	// Wait for all goroutines
	waitGroup.Wait()

	// Return the result array, resultAmount
	return rsc.ResultArray
}

// The SliceContains() function returns whether or not the provided
// slice contains the provided map
func SliceContainsMap(slice []map[string]string, _map map[string]string) bool {
	// Iterate over the slice
	for i := 0; i < len(slice); i++ {
		// if the slice value equals the string then return true
		if fmt.Sprint(slice[i]) == fmt.Sprint(_map) {
			return true
		}
	}
	// Else return false
	return false
}
