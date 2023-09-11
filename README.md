# Unicorn Builder API
This API allows users to create and retrieve unicorns using a storage and real-time creation system.
`Unicorn` is a type created in golang with 3 attributes: name, id, and capabilities. All three are created randomly according to a set of files listed on the `repository` package, which is going to be mentioned after 

```go
type Unicorn struct {
	Name         string   `json:"name"`
	Id           int      `json:"id"`
	Capabilities []string `json:"capabilities"`
}
```
## API Endpoints

### `/api/get-unicorn` (GET)
This endpoint allows users to request unicorns. Users can specify the number of unicorns they want and receive a JSON response containing the requested number.
#### Parameters
- `amount` (integer): The number of unicorns to retrieve. Established as a query param

#### Example Request
```http
GET /api/get-unicorn?amount=3
```
#### Example Response
```json
[
    {
        "name": "pointless-retha",
        "id": 1,
        "capabilities": [
            "code",
            "change color",
            "walk"
        ]
    },
    {
        "name": "genuine-korey",
        "id": 2,
        "capabilities": [
            "fulfill wishes",
            "lazy",
            "swim"
        ]
    },
    {
        "name": "fruitful-trenton",
        "id": 3,
        "capabilities": [
            "sing",
            "cry",
            "talk"
        ]
    }
]
```

## Data generation process
1. When a user sends a request to the API, they provide a query parameter called "amount," which tells the program how many unicorns they need.

2. The API first checks if there are enough unicorns in the *store*. This is a data structure that acts as a Last-In-First-Out (LIFO) stack and stores pre-created unicorns ready for immediate use.

3. The API uses those unicorns to finish the response if there are sufficient unicorns in the *store* to satisfy the user's request. These unicorns are selected following the LIFO pattern, meaning the most recently created unicorns are used first

4. If there aren't enough unicorns in the *store* to meet the user's request or if the *store* is empty, the API uses the available unicorns and then creates the remaining in real-time to fulfill the user's request. These real-time unicorns are created sequentially and follow the First-In-First-Out (FIFO) pattern, meaning the newly created unicorns are used first.

5. During the real-time unicorn creation process, the API displays the *request-id* of the unicorn being generated. Once fully created, the *request-id* is displayed again, allowing the user to keep track of the unicorns specifically generated to fulfill their request.

This approach ensures that unicorns stored in the *store* are used before creating new unicorns in real-time, and it follows a consistent logic for fulfilling user requests. Real-time unicorns are generated sequentially and used in the order they were created to ensure fair and predictable distribution.

### Example of process
``` cmd
$ go run cmd/main.go 
2023/09/11 14:56:04 Listening at port: 8888
- Unicorn with id 1 was added to the store
Processing new request...
- Unicorn with id 1 was taken from the store
- Creating unicorn with id 2...
- Unicorn with id 2 was successfully created
- Creating unicorn with id 3...
- Unicorn with id 3 was successfully created
Request has been completed.
```
With this example can be seen that *unicorn 1* was stored in the *store* and when a request came in, it was immediately pushed out to be used in the response. Afterward, the *unicorn 2* and *unicorn 3* were live-created and added to the response

## Folder structure
The Unicorn Builder API is organized into the following components:

* Handlers: Responsible for handling incoming HTTP requests, extracting data, and invoking the appropriate services.
* Services: Implement the core logic of the API, including creating and retrieving unicorns.
* Stack: Manages a stack of unicorns that can be used to fulfill requests.
* Repository: Handles data retrieval, such as fetching names and adjectives for unicorn creation.
* Util: Contains utility functions for generating random names, capabilities, and IDs.




