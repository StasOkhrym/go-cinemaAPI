package handlers

import (
	"cinemaAPI/db"
	"cinemaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	db.DB.Find(&movies)
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func GetMovie(c *gin.Context) {
	var movie models.Movie
	if err := db.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func CreateMovie(c *gin.Context) {
	var input models.CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := models.Movie{Title: input.Title, Description: input.Description, Duration: input.Duration, Image: input.Image}
	db.DB.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func UpdateMovie(c *gin.Context) {
	var movie models.Movie
	if err := db.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Model(&movie).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func DeleteMovie(c *gin.Context) {
	var movie models.Movie
	if err := db.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.DB.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
