# Building a Simple API using Go

This is a simple API for building an API using Gin. 

## Prerequisites

* Go version go1.22.1

## Getting Started.

1. Run the command `go get .`
2. Open up your browser or Thunderclient or Postman and visit `https:127.0.0:8080/albums`, an endpoint to the albums.

## API Endpoints Available.

### 1. Albums

This is an endpoint for retrieving all the albums
```
GET /albums
```

### Responses
```json
{
    "ID": "string",
    "Title": "string",
    "Artist": "string",
    "Price": "float64"
},
```

### Status Codes
200: Successful

## Author
* Joel Nyongesa