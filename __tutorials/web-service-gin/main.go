// A standalone program (as opposed to a library) is always in package main
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main()  {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)

	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}

/*
* gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more
*/
func getAlbums(context *gin.Context)  {
	/*
	* Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	The function’s first argument is the HTTP status code you want to send to the client.
	
	Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.z

	you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
	*/
	context.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	// Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func addAlbum(context *gin.Context)  {
	var newAlbum album

	// BindJSON binds the received JSON to
    // newAlbum
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	// Add a 201 status code to the response, along with JSON representing the album you added
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

