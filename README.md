To run this, you just need to have Go installed and type "go run main.go"

This will start a server listening on localhost:8081 by default. 
You can make any kind of call to it and have the middleware chain validate and process it.
If all the checks pass and the transforms are successful, the response returned is the request url, body, and headers wrapped up in a JSON
