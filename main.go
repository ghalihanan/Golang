package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "Blue Eyes", Artist: "Michael", Price: 24.99},
	{ID: 2, Title: "Kiss Me", Artist: "Arnold", Price: 24.99},
	{ID: 3, Title: "In the sky", Artist: "Julian", Price: 24.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
