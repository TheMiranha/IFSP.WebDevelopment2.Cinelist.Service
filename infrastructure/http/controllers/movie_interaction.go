package controllers

import (
	"cinelist/application/usecases"
	"cinelist/domain/dtos"
	infrastructure_utils "cinelist/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieInteractionController struct {
	usecase usecases.MovieInteractionUseCase
}

func NewMovieInteractionController(usecase usecases.MovieInteractionUseCase) MovieInteractionController {
	return MovieInteractionController{
		usecase: usecase,
	}
}

func (c *MovieInteractionController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/movies/favorite", c.FavoriteMovie)
	router.POST("/movies/to-watch", c.ToWatchMovie)
	router.POST("/movies/watched", c.CreateWatched)
}

func (c *MovieInteractionController) FavoriteMovie(ctx *gin.Context) {
	idFromCtx, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuário não encontrado no token"})
		return
	}

	userID := idFromCtx.(uuid.UUID)

	var payload dtos.FavoriteMovieDTO
	err := ctx.BindJSON(&payload)
	if err != nil {
		infrastructure_utils.ThrowInvalidRequest(ctx)
		return
	}

	movieID, err := uuid.Parse(payload.MovieId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewRequestError("Invalid movie ID"))
		return
	}

	isFavorited, errResponse := c.usecase.FavoriteMovie(userID, movieID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	var message string
	if isFavorited {
		message = "Movie favorited successfully"
	} else {
		message = "Movie unfavorited successfully"
	}

	ctx.JSON(http.StatusOK, dtos.SuccessResponseDTO{
		Success: true,
		Message: message,
	})
}

func (c *MovieInteractionController) ToWatchMovie(ctx *gin.Context) {
	idFromCtx, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuário não encontrado no token"})
		return
	}

	userID := idFromCtx.(uuid.UUID)

	var payload dtos.ToWatchMovieDTO
	err := ctx.BindJSON(&payload)
	if err != nil {
		infrastructure_utils.ThrowInvalidRequest(ctx)
		return
	}

	movieID, err := uuid.Parse(payload.MovieId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewRequestError("Invalid movie ID"))
		return
	}

	isToWacth, errResponse := c.usecase.ToWatchMovie(userID, movieID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	var message string
	if isToWacth {
		message = "Movie added to toWatch list successfully"
	} else {
		message = "Movie removed from toWatch list successfully"
	}

	ctx.JSON(http.StatusOK, dtos.SuccessResponseDTO{
		Success: true,
		Message: message,
	})
}

func (c *MovieInteractionController) CreateWatched(ctx *gin.Context) {
	idFromCtx, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuário não encontrado no token"})
		return
	}

	userID := idFromCtx.(uuid.UUID)

	var payload dtos.WatchedMovieDTO
	err := ctx.BindJSON(&payload)
	if err != nil {
		infrastructure_utils.ThrowInvalidRequest(ctx)
		return
	}

	movieID, err := uuid.Parse(payload.MovieId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewRequestError("Invalid movie ID"))
		return
	}

	errResponse := c.usecase.CreateWatched(userID, movieID, payload.Rate, payload.Description)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, dtos.SuccessResponseDTO{
		Success: true,
		Message: "Watched entry created successfully",
	})
}
