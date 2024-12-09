package controller

import (
	"net/http"
	"savioafs/daily-diet-app-go/internal/dto"
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MealController struct {
	MealUseCase usecase.MealUsecase
}

func NewMealController(mealUsecase usecase.MealUsecase) *MealController {
	return &MealController{MealUseCase: mealUsecase}
}

func (c *MealController) Create(ctx *gin.Context) {
	var mealInput dto.MealInputDTO
	err := ctx.BindJSON(&mealInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "invalid input",
		})
		return
	}

	meal, err := entity.NewMeal(mealInput.UserID, mealInput.Name, mealInput.Description, mealInput.Date, mealInput.IsDiet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

	meal.Validate()

	createdMeal, err := c.MealUseCase.Create(meal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "could not create meal",
		})
		return
	}

	mealOutput := dto.MealOutputDTO{
		ID:          createdMeal.ID,
		UserID:      createdMeal.UserID,
		Name:        createdMeal.Name,
		Description: createdMeal.Description,
		Date:        createdMeal.Date,
		IsDiet:      createdMeal.IsDiet,
	}

	ctx.JSON(http.StatusCreated, mealOutput)
}

func (c *MealController) GetMealByID(ctx *gin.Context) {
	id := ctx.Param("id")

	meal, err := c.MealUseCase.FindMealByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

	mealOutput := dto.MealOutputDTO{
		ID:          meal.ID,
		UserID:      meal.UserID,
		Name:        meal.Name,
		Description: meal.Description,
		Date:        meal.Date,
		IsDiet:      meal.IsDiet,
	}

	ctx.JSON(http.StatusOK, mealOutput)
}
