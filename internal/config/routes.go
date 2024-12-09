package config

import (
	"database/sql"
	"savioafs/daily-diet-app-go/internal/controller"
	"savioafs/daily-diet-app-go/internal/repository"
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(dbConn *sql.DB) *gin.Engine {
	server := gin.Default()

	mealRepository := repository.NewMealRepositoryPG(dbConn)
	mealUseCase := usecase.NewMealUseCase(mealRepository)
	mealController := controller.NewMealController(mealUseCase)

	mealsGroup := server.Group("/meals")
	{
		mealsGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		mealsGroup.GET("", mealController.Create)
	}

	return server
}
