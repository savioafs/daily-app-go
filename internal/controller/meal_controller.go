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
	userID := ctx.DefaultQuery("user_id", "")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user_id is required",
		})
		return
	}

	err := ctx.BindJSON(&mealInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "invalid input",
		})
		return
	}

	meal, err := entity.NewMeal(userID, mealInput.Name, mealInput.Description, mealInput.Date, mealInput.IsDiet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

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

func (c *MealController) GetAllMealsByUser(ctx *gin.Context) {
	userID := ctx.DefaultQuery("user_id", "")

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "user id is required",
		})
		return
	}

	var mealsOutput []dto.MealOutputDTO

	meals, err := c.MealUseCase.GetAllMealsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

	if len(meals) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message:": "no meals found for user",
		})
	}

	for _, meal := range meals {

		mealOutput := dto.MealOutputDTO{
			ID:          meal.ID,
			UserID:      meal.UserID,
			Name:        meal.Name,
			Description: meal.Description,
			Date:        meal.Date,
			IsDiet:      meal.IsDiet,
		}

		mealsOutput = append(mealsOutput, mealOutput)
	}

	ctx.JSON(http.StatusOK, mealsOutput)
}

func (c *MealController) GetMealsUserByStatus(ctx *gin.Context) {
	userID := ctx.DefaultQuery("user_id", "")
	status := ctx.DefaultQuery("status", "false")

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "user id is required",
		})
		return
	}

	var boolStatus bool
	if status == "true" {
		boolStatus = true
	}

	meals, err := c.MealUseCase.GetMealsUserByStatus(userID, boolStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "could not retrieve meals",
		})
		return
	}

	ctx.JSON(http.StatusOK, meals)
}

func (c *MealController) MetricsMealsByUser(ctx *gin.Context) {
	userID := ctx.DefaultQuery("user_id", "")

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "user id is required",
		})
		return
	}

	metrics, err := c.MealUseCase.MetricsMealsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "cannot get metrics by user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_meals":          metrics["total_meals"],
		"total_meals_diet":     metrics["total_meals_diet"],
		"total_meals_non_diet": metrics["total_meals_non_diet"],
		"diet_percent":         metrics["diet_percent"],
		"non_diet_percent":     metrics["non_diet_percent"],
	})
}

func (c *MealController) UpdateMeal(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "id is required",
		})
		return
	}

	var meal entity.Meal

	err := ctx.BindJSON(&meal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "invalid input",
		})
		return
	}

	err = c.MealUseCase.UpdateMeal(id, &meal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message:": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully",
	})
}
