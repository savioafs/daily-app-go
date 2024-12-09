package main

import (
	"fmt"
	"savioafs/daily-diet-app-go/internal/controller"
	"savioafs/daily-diet-app-go/internal/db"
	"savioafs/daily-diet-app-go/internal/repository"
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	mealRepository := repository.NewMealRepositoryPG(dbConn)
	mealUseCase := usecase.NewMealUseCase(mealRepository)
	mealController := controller.NewMealController(mealUseCase)

	fmt.Println(mealController)

	server.Run(":8080")
}
