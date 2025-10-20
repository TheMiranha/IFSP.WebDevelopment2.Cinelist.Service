package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActorController struct {
	usecase usecases.ActorUseCase
}

func NewActorController(usecase usecases.ActorUseCase) ActorController {
	return ActorController{
		usecase: usecase,
	}
}

func (uc *ActorController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/actors", uc.GetAll)
}

func (uc *ActorController) GetAll(ctx *gin.Context) {
	actors, err := uc.usecase.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.ActorsResponseDTO{
		Success: true,
		Data:    actors,
	})
}
