package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)


// Album that represents data about a record album
type album struct{
	ID string `json: "id"`
	Title string `json: "title"`
	Artist string `json: "artist"`
	Price float64 `json: "price"`
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// Albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Midnight Train", Artist: "Sauti Sol", Price: 10.00},
	{ID: "2", Title: "Time", Artist: "H_art the Band", Price: 20.00},
	{ID: "3", Title: "Alusa Why Are You Topless", Artist: "Bien", Price: 25.00},
}

// getAlbums (GET) method
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}