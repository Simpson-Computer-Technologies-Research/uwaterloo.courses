package redis

// Import modules
import (
	"context"

	"github.com/go-redis/redis/v9"
)

/*
	RUN docker run --name redis-test-instance -p 6379:6379 -d redis

	The redis database is primarily used for testing without wifi or
	for the redisearch full text search (which is currently unfinished).

	I decided to go with an in memory map that holds the course data because
	it's 3-10x faster than using the redis database. Since the university of
	waterloo doesn't hold a lot of course data, using the in memory map
	won't overflow the api's memory usage.

	The in memory map and it's functions can be found inside
	the ./cache/cache.go file.

	The in memory cache is reset everytime the api scrapes the courses
	from the official university of waterloo website. Same goes for
	refreshing the redis database.

*/

// Define Variables
var (
	// Context: context.Context -> Used for the Redis Cache
	Context context.Context = context.Background()

	// RedisCache: *redis.Client -> Used to Connect to the redis database
	RedisCache *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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
