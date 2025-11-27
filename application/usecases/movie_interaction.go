package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	"cinelist/infrastructure/database/repositories"

	"github.com/google/uuid"
)

type MovieInteractionUseCase struct {
	repo repositories.MovieInteractionRepository
}

func NewMovieInteractionUseCase(repo repositories.MovieInteractionRepository) MovieInteractionUseCase {
	return MovieInteractionUseCase{repo: repo}
}

func (uc *MovieInteractionUseCase) FavoriteMovie(userID uuid.UUID, movieID uuid.UUID) *dtos.RequestError {
	err := uc.repo.CreateFavorite(userID, movieID)
	if err != nil {
		return dtos.NewRequestError("Error while favoriting movie")
	}
	return nil
}

func (uc *MovieInteractionUseCase) AddToWatch(userID uuid.UUID, movieID uuid.UUID) *dtos.RequestError {
	err := uc.repo.CreateToWatch(userID, movieID)
	if err != nil {
		return dtos.NewRequestError("Error while adding movie to watch list")
	}
	return nil
}

func (uc *MovieInteractionUseCase) CreateWatched(userID uuid.UUID, movieID uuid.UUID, rate float64, description string) *dtos.RequestError {
	watched := entities.Watched{
		User:        userID,
		Movie:       movieID,
		Rate:        rate,
		Description: description,
	}

	err := uc.repo.CreateWatched(watched)
	if err != nil {
		return dtos.NewRequestError("Error while creating watched entry")
	}
	return nil
}

