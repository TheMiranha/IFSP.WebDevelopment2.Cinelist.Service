package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	usecase usecases.MovieUseCase
}

func NewMovieController(usecase usecases.MovieUseCase) MovieController {
	return MovieController{
		usecase: usecase,
	}
}

func (uc *MovieController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/movies", uc.GetAll)
}

func (uc *MovieController) GetAll(ctx *gin.Context) {
	movies, err := uc.usecase.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.MoviesResponseDTO{
		Success: true,
		Data:    movies,
	})
}
