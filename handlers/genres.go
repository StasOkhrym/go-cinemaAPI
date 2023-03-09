package handlers

import (
	"cinemaAPI/db"
	"cinemaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	db.DB.Find(&genres)
	c.JSON(http.StatusOK, gin.H{"data": genres})
}

func GetGenre(c *gin.Context) {
	var genre models.Genre
	if err := db.DB.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func CreateGenre(c *gin.Context) {
	var input models.CreateGenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre := models.Genre{Name: input.Name}
	db.DB.Create(&genre)
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func UpdateGenre(c *gin.Context) {
	var genre models.Genre
	if err := db.DB.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateGenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Model(&genre).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func DeleteGenre(c *gin.Context) {
	var genre models.Genre
	if err := db.DB.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.DB.Delete(&genre)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
