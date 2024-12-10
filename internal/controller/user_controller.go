package controller

import (
	"fmt"
	"net/http"
	"savioafs/daily-diet-app-go/internal/dto"
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

type UserController struct {
	UserUseCase  usecase.UserUseCase
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserController(usecase usecase.UserUseCase, jwt *jwtauth.JWTAuth, expiresJwt int) *UserController {
	return &UserController{
		UserUseCase:  usecase,
		Jwt:          jwt,
		JwtExpiresIn: expiresJwt,
	}
}

func (c *UserController) GetJWT(ctx *gin.Context) {
	var user dto.GetJWTInput
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message:": err.Error(),
		})
		return
	}

	userFind, err := c.UserUseCase.FindByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	if !userFind.ValidatePassword(user.Password) {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	_, tokenString, _ := c.Jwt.Encode(map[string]interface{}{
		"sub": userFind.ID,
		"exp": time.Now().Add(time.Second * time.Duration(c.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	ctx.JSON(http.StatusOK, accessToken)
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

func (c *UserController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	fmt.Println("entrou aqui")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email is required",
		})
		return
	}

	user, err := c.UserUseCase.FindByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userOutput := dto.UserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	fmt.Println(userOutput)

	ctx.JSON(http.StatusOK, userOutput)
}
