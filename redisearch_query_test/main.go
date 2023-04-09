package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
)

// Course struct
type Course struct {
	score          float32
	docId          string
	id             string
	name           string
	title          string
	description    string
	pre_requisites string
	components     string
	unit           string
}

func main() {
	RedisAPISearch()
	// RediSearchAPI()

	// run the program: go run main.go
}

// InitSchema initializes the schema for insetring data into the redis db
func InitSchema(client *redisearch.Client) {
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
}

// Insert data into the redis db
func InsertData(client *redisearch.Client, course Course) {
	// Create a new document that's going to be used to insert data into the db
	var doc = redisearch.NewDocument(course.docId, course.score)
	doc.Set("components", course.components).
		Set("description", course.description).
		Set("id", course.id).
		Set("name", course.name).
		Set("pre_requisites", course.pre_requisites).
		Set("title", course.title).
		Set("unit", course.unit)

	// Update the database, insert the document
	if err := client.Index([]redisearch.Document{doc}...); err != nil {
		fmt.Println(err)
	}
}

// Perform a search using the redis/v8 package
func RediSearchAPI() {
	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	var client *redisearch.Client = redisearch.NewClient("localhost:6379", "index")

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

// Redis API Search
func RedisAPISearch() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:7501"})
	ctx := context.Background()

	// Track time
	var startTime time.Time = time.Now()

	// Search for the term "math" and return the first 100 results
	var values, _ = rdb.Do(ctx,
		"FT.SEARCH", "courses", "math",
		"RETURN", "0", "LIMIT", "0", "100",
	).Slice()

	// Print the time it took to run the query
	fmt.Println(time.Since(startTime))

	// Print the document ids
	for _, id := range values {
		fmt.Println(id)
	}
}
