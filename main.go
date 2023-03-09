package main

import (
	"cinemaAPI/db"
	"cinemaAPI/handlers"
	"cinemaAPI/models"

	"github.com/gin-gonic/gin"
)

func startDB() {
	db.Init()
	db.DB.AutoMigrate(&models.Movie{}, &models.Genre{}, &models.Actor{})
}

func main() {
	startDB()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, world!"})
	})

	v1 := router.Group("/api/v1")
	{
		cinema := v1.Group("/cinema")
		{
			movies := cinema.Group("/movies")
			{
				movies.GET("/", handlers.GetMovies)
				movies.GET("/:id", handlers.GetMovie)
				movies.POST("/", handlers.CreateMovie)
				movies.PUT("/:id", handlers.UpdateMovie)
				movies.DELETE("/:id", handlers.DeleteMovie)
			}
			genres := cinema.Group("/genres")
			{
				genres.GET("/", handlers.GetGenres)
				genres.GET("/:id", handlers.GetGenre)
				genres.POST("/", handlers.CreateGenre)
				genres.PUT("/:id", handlers.UpdateGenre)
				genres.DELETE("/:id", handlers.DeleteGenre)
			}
		}
	}

	router.Run()
}
