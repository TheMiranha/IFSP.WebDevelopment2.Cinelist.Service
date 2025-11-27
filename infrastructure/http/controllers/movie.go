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
	router.GET("/movies/search", uc.Search)
}

func (uc *MovieController) GetAll(ctx *gin.Context) {
	term := ctx.Query("term")

	if term != "" {
		movies, err := uc.usecase.GetAllByTitle(term)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, dtos.MoviesResponseDTO{
			Success: true,
			Data:    movies,
		})
		return
	}

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

func (uc *MovieController) Search(ctx *gin.Context) {
	term := ctx.Query("term")

	if term == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "term query parameter is required",
		})
		return
	}

	moviesWithCast, err := uc.usecase.Search(term)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.MoviesWithCastResponseDTO{
		Success: true,
		Data:    moviesWithCast,
	})
}
