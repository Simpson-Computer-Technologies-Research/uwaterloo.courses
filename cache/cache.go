package cache

// Import modules
import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Hold Course data in memory cache
var Cache [][]byte

// The Set() function add the data to the
// cache byte slice
func Set(value map[string]string) {
	var tmp, _ = json.Marshal(value)
	Cache = append(Cache, tmp)
}

// The GetCourses() function iterates over the cache
// slice and checks if the data contains the provided
// query/subject
func GetCourses(query []byte, subject []byte) []byte {
	subject = []byte(fmt.Sprintf(`,"title":"%s `, subject))
	query = bytes.ToLower(query)

	// Create result slices
	var (
		subResult   []byte = []byte{}
		queryResult []byte = []byte{}
	)
	// Iterate over the cache
	for i := 0; i < len(Cache); i++ {
		if len(queryResult) > 400000 {
			break
		}
		// Check if the cache contains the subject
		if bytes.Contains(Cache[i], subject) {
			subResult = append(subResult, append(Cache[i], ',')...)
		} else

		// Check if the lowercase cache contains the subject
		if bytes.Contains(bytes.ToLower(Cache[i]), query) {
			queryResult = append(queryResult, append(Cache[i], ',')...)
		}
	}
	var res []byte = append(subResult, queryResult...)
	if len(res) > 0 {
		return res
	}
	return []byte{'{', '}'}
}
