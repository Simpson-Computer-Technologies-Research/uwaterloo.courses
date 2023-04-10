from redis import ResponseError
import redisearch, time

# // Initialize the client function
def initialize_client():
    SCHEMA = (
        redisearch.TextField("components"),
        redisearch.TextField("description"),
        redisearch.TextField("id"),
        redisearch.TextField("name"),
        redisearch.TextField("pre_requisites"),
        redisearch.TextField("title"),
        redisearch.TextField("unit"),
    )
    client: redisearch.Client = redisearch.Client("course", 
        host = "localhost", 
        port = 6379
    )
    definition: redisearch.IndexDefinition = redisearch.IndexDefinition(
        prefix = [':course.Course:']
    )
    try: client.info()
    except ResponseError: client.create_index(SCHEMA, definition = definition)
    return client

# // Initialize the client
client: redisearch.Client = initialize_client()

# // Find by description
def find(client: redisearch.Client):
    start_time = time.time()
    res = client.search("CS*")
    for result in res.docs:
        print(result)
    print("Search took: " + str(time.time() - start_time) + " seconds")

# // Call find
find(client = client)