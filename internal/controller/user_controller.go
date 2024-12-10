package controller

import (
	"net/http"
	"savioafs/daily-diet-app-go/internal/dto"
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) *UserController {
	return &UserController{UserUseCase: usecase}
}

func (c *UserController) Create(ctx *gin.Context) {
	var userInput dto.UserInputDTO

	err := ctx.BindJSON(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": "invalid data to create user",
		})
		return
	}

	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

	err = c.UserUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message:": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message:": "created successfully",
	})
}
