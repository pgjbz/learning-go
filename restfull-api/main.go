package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pgjbz.dev/restfull-api/db"
	"pgjbz.dev/restfull-api/models"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAlbums())
}

func postAlbums(c *gin.Context) {
	var (
		newAlbum models.Album
	)

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	db.AddAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.ParseUint(stringId, 0, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "cannot parse string to integer"})
		return
	}
	for _, album := range db.GetAlbums() {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
