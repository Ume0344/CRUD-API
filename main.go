package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Creating a struct to define album
// The content inside backticks ` ` are tags which describes how the struct will be
// encoded when converting to other format. Here, we are going to convert it to
// json that's why the json tag
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artists"`
	Price  float64 `json:"price"`
}

// An array/slice to define some albums
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Whenever /albums endpoint is called, getAlbums function is invoked
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {
	// First argument to IntendedJSON is the status you want to send to client
	// and second is the struct you want to serialize

	// Serialize/converting the struct into json and add it to response
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
