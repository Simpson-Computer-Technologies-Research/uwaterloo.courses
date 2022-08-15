package redis

import "fmt"

// The CreateSchema() function is used to create the
// redis cache full text search schema.
// This schema is created for the redis cache full text
// search mechanism
func CreateSchema() {
	RedisCache.Do(Context,
		"FT.CREATE", "courses", "ON", "JSON", "PREFIX", "1", "course:", "NOOFFSETS",
		"SCHEMA",
		"$.title", "AS", "title", "TEXT",
		"$.description", "AS", "description", "TEXT",
		"$.name", "AS", "name", "TEXT",
		"$.id", "AS", "id", "TEXT",
		"$.components", "AS", "components", "TEXT",
		"$.unit", "AS", "unit", "TEXT",
		"$.anti_reqs", "AS", "anti_reqs", "TEXT",
		"$.pre_reqs", "AS", "pre_reqs", "TEXT",
	)
}

// The QuerySearch() function is used to perform
// a full text search for the given query string
func QuerySearch(query string) {
	var res, _ = RedisCache.Do(Context,
		"FT.SEARCH", "courses", query,
		"RETURN", "0", "LIMIT", "0", "100",
	).Slice()
	fmt.Println(res)
}
