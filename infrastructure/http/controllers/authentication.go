package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
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
}

func (c *AuthenticationController) SignUp(ctx *gin.Context) {
	var payload dtos.SignUpDTO
	err := ctx.BindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewRequestError("Invalid request"))

		return
	}
	response, usecaseError := c.usecase.Create(payload)
	if usecaseError != nil {
		ctx.JSON(http.StatusInternalServerError, usecaseError)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
