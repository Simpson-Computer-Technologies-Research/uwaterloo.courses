package main

import (
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func main() {
	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	var client *redisearch.Client = redisearch.NewClient("localhost:6379", "index")

	// Create a schema
	var schema *redisearch.Schema = redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("components")).
		AddField(redisearch.NewTextField("description")).
		AddField(redisearch.NewTextField("id")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("pre_requisites")).
		AddField(redisearch.NewTextField("title")).
		AddField(redisearch.NewTextField("unit"))

	// Create the index with the given schema
	if err := client.CreateIndex(schema); err != nil {
		fmt.Println(err)
	}

	// Track the start time
	var startTime time.Time = time.Now()

	// Searching with limit and sorting
	var docs, _, _ = client.Search(redisearch.NewQuery("math").Limit(0, 600))

	// Print the first document id, title, total results and error
	fmt.Println(docs)

	// Print the time it took to run the query
	// The average time is about 60ms which is way slower than my uwaterloo.courses query method
	fmt.Println(time.Since(startTime))

	// run the program: go run main.go
}
