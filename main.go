package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue train", Artist: "Jhon Coltrane", Price: 56.99}, 
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}, 
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 36.99}, 	
}


func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err)
		return 
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID (c *gin.Context){
	id := c.Param("id")
	for _, a := range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main () {
	r := gin.Default()

	//GET ALBUMS
	r.GET("/albums", getAlbums)

	//GET BY ID
	r.GET("/albums/:id", getAlbumByID)

	//POST ALBUMS
	r.POST("/albums", postAlbums)

	r.Run("localhost:8000")
}


