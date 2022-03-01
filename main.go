package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Staying' Alive", Artist: "The Bee Gees", Price: 56.99},
	{ID: "2", Title: "Sexual Healing", Artist: "Marvin Gaye", Price: 49.99},
	{ID: "3", Title: "Disco Inferno", Artist: "The Trammps", Price: 52.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbum locates the album whose ID matches the id.
func getAlbum(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of items for the specific album id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found :("})
}

// postAlbums adds an album from JSON received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
