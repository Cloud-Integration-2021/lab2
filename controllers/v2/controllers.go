package v2

import (
	"github.com/gin-gonic/gin"
	"lab2/models"
	"net/http"
)

// FindActorsByMovieId
// GET /actors/:idMovie
func FindActorsByMovieId(c *gin.Context) {
	c.JSON(http.StatusOK, models.GenerateRandomActors())
}
