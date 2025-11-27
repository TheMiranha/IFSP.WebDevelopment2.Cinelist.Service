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
	router.POST("/movies/to-watch", c.AddToWatch)
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

	errResponse := c.usecase.FavoriteMovie(userID, movieID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, dtos.SuccessResponseDTO{
		Success: true,
		Message: "Movie favorited successfully",
	})
}

func (c *MovieInteractionController) AddToWatch(ctx *gin.Context) {
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

	errResponse := c.usecase.AddToWatch(userID, movieID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, dtos.SuccessResponseDTO{
		Success: true,
		Message: "Movie added to watch list successfully",
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

