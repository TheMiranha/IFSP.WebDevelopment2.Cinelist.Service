package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	usecase usecases.UserUseCase
}

func NewUserController(usecase usecases.UserUseCase) UserController {
	return UserController{usecase: usecase}
}

func (c *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/user/me", c.MeRoute)
}

func (c *UserController) MeRoute(ctx *gin.Context) {
	idFromCtx, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuário não encontrado no token"})
		return
	}

	fmt.Println(idFromCtx)

	userID := idFromCtx.(uuid.UUID)

	userData, errResponse := c.usecase.GetUserById(userID)

	if errResponse != nil {
		ctx.JSON(http.StatusNotFound, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, dtos.UserResponseDTO{
		Success: true,
		Data:    userData,
	})
}
