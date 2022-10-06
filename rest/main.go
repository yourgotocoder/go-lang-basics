package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Mayer", Artist: "John Mayer", Price: 6.99},
	{ID: "2", Title: "New Light", Artist: "John Mayer", Price: 8.99},
	{ID: "3", Title: "Room of Squares", Artist: "John Mayer", Price: 9.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:3000")
}
