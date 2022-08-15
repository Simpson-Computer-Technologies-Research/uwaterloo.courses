package redis

// Import modules
import (
	"context"

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

// The Contains() function checks whether
// or not the redis cache contains a key
func Contains(key string) bool {
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
