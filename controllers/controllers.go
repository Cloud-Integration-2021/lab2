package controllers

import (
	"net/http"

	"lab2/models"

	"github.com/gin-gonic/gin"
)


// FindActorsByMovieId
// GET /actors/:idMovie
func FindActorsByMovieId(c *gin.Context) {
	c.JSON(http.StatusOK, models.GenerateRandomActors())
}