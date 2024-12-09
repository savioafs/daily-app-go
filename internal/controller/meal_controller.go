package controller

import (
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MealController struct {
	MealUseCase usecase.MealUsecase
}

func NewMealController(mealUsecase usecase.MealUsecase) *MealController {
	return &MealController{MealUseCase: mealUsecase}
}

func (c *MealController) Create(ctx *gin.Context) {}
