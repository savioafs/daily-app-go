package config

import (
	"database/sql"
	"savioafs/daily-diet-app-go/internal/controller"
	"savioafs/daily-diet-app-go/internal/middleware"
	"savioafs/daily-diet-app-go/internal/repository"
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

func SetupRoutes(dbConn *sql.DB, expiresIn int, jwtAuth *jwtauth.JWTAuth) *gin.Engine {
	server := gin.Default()

	mealRepository := repository.NewMealRepositoryPG(dbConn)
	mealUseCase := usecase.NewMealUseCase(mealRepository)
	mealController := controller.NewMealController(mealUseCase)

	mealsGroup := server.Group("/meals")
	{
		mealsGroup.Use(middleware.JWTAuthMiddleware(jwtAuth))

		mealsGroup.POST("", mealController.Create)
		mealsGroup.GET("/:id", mealController.GetMealByID)
		mealsGroup.GET("", mealController.GetAllMealsByUser)
		mealsGroup.GET("/status", mealController.GetMealsUserByStatus)
		mealsGroup.GET("/metrics", mealController.MetricsMealsByUser)
		mealsGroup.PUT("/:id", mealController.UpdateMeal)
	}

	userRepository := repository.NewUserRepositoryPG(dbConn)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUsecase, jwtAuth, expiresIn)

	usersGroup := server.Group("/users")
	{
		usersGroup.POST("", userController.Create)
		usersGroup.POST("/generate_token", userController.GetJWT)
		usersGroup.GET("/:email", userController.FindByEmail)
	}

	return server
}
