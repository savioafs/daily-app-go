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
		mealsGroup.POST("", mealController.Create)
		mealsGroup.GET("/:id", mealController.GetMealByID)
		mealsGroup.GET("", mealController.GetAllMealsByUser)
		mealsGroup.GET("/status", mealController.GetMealsUserByStatus)
		mealsGroup.GET("/metrics", mealController.MetricsMealsByUser)
		mealsGroup.PUT("/:id", mealController.UpdateMeal)
	}

	userRepository := repository.NewUserRepositoryPG(dbConn)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUsecase)

	usersGroup := server.Group("/users")
	{
		usersGroup.POST("", userController.Create)
	}

	return server
}
