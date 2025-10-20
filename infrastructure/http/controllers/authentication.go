package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
	infrastructure_utils "cinelist/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	usecase usecases.AuthenticationUseCase
}

func NewAuthenticationController(usecase usecases.AuthenticationUseCase) AuthenticationController {
	return AuthenticationController{usecase: usecase}
}

func (c *AuthenticationController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/auth/sign-up", c.SignUp)
	router.POST("/auth/sign-in", c.SignIn)
}

func (c *AuthenticationController) SignIn(ctx *gin.Context) {
	var payload dtos.SignInDTO

	err := ctx.BindJSON(&payload)

	if err != nil {
		infrastructure_utils.ThrowInvalidRequest(ctx)
		return
	}

	usecaseResponse, usecaseError := c.usecase.SignIn(payload)

	if usecaseError != nil {
		ctx.JSON(http.StatusUnauthorized, usecaseError)
		return
	}

	ctx.JSON(http.StatusOK, usecaseResponse)
}

func (c *AuthenticationController) SignUp(ctx *gin.Context) {
	var payload dtos.SignUpDTO
	err := ctx.BindJSON(&payload)

	if err != nil {
		infrastructure_utils.ThrowInvalidRequest(ctx)
		return
	}
	response, usecaseError := c.usecase.Create(payload)
	if usecaseError != nil {
		ctx.JSON(http.StatusInternalServerError, usecaseError)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
