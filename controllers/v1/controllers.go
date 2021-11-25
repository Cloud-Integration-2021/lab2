package v1

import (
	"net/http"

	"lab2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type CreateMovieInput struct {
	Title       string `json:"title" binding:"required"`
	ReleaseDate string `json:"releaseDate" binding:"required"`
	Plot        string `json:"plot"`
}

type UpdateMovieInput struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Plot        string `json:"plot"`
}

// FindMovies
// GET /Movies
func (DB *Database) FindMovies(c *gin.Context) {
	var Movies []models.Movie
	DB.Find(&Movies)

	c.JSON(http.StatusOK, Movies)
}

// FindMovieById
// GET /Movies/:id
func (DB *Database) FindMovieById(c *gin.Context) {
	// Get model if exist
	var Movie models.Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&Movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, Movie)
}

// CreateMovie
// POST /Movies
func (DB *Database) CreateMovie(c *gin.Context) {
	// Validate input
	var input CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Movie
	Movie := models.Movie{Title: input.Title, ReleaseDate: input.ReleaseDate, Plot: input.Plot}
	DB.Create(&Movie)

	c.JSON(http.StatusOK, Movie)
}

// UpdateMovie
// PUT /Movies/:id
func (DB *Database) UpdateMovie(c *gin.Context) {
	// Get model if exist
	var movie models.Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.Title = input.Title
	movie.ReleaseDate = input.ReleaseDate
	movie.Plot = input.Plot

	DB.Save(&movie)

	c.JSON(http.StatusOK, movie)
}

// DeleteMovie
// DELETE /Movies/:id
func (DB *Database) DeleteMovie(c *gin.Context) {
	// Get model if exist
	var Movie models.Movie
	if err := DB.Where("id = ?", c.Param("id")).First(&Movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&Movie)

	c.JSON(http.StatusOK, Movie)
}