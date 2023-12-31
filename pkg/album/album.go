package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @Summary albums example
// @Schemes
// @Description albums test
// @Tags albums
// @Success 200 {string} GetAlbums
// @Router /api/v1/albums [get]
// getAlbums responds with the list of all albums as JSON.
// @Security BearerAuth
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
