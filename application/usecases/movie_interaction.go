package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	domain_repositories "cinelist/domain/repositories"

	"github.com/google/uuid"
)

type MovieInteractionUseCase struct {
	repo domain_repositories.MovieInteractionRepository
}

func NewMovieInteractionUseCase(repo domain_repositories.MovieInteractionRepository) MovieInteractionUseCase {
	return MovieInteractionUseCase{repo: repo}
}

func (uc *MovieInteractionUseCase) FavoriteMovie(userID uuid.UUID, movieID uuid.UUID) (bool, *dtos.RequestError) {
	// Verifica se o filme já está favoritado
	favorite, err := uc.repo.GetFavoriteByUserAndMovie(userID, movieID)
	if err == nil {
		// Verifica se a entidade não está vazia (encontrou)
		if favorite.User != uuid.Nil && favorite.Movie != uuid.Nil {
			// Já está favoritado, então remove
			err = uc.repo.DeleteFavorite(userID, movieID)
			if err != nil {
				return false, dtos.NewRequestError("Error while unfavoriting movie")
			}
			return false, nil // false = foi removido
		}
	}

	// Não está favoritado, então adiciona
	err = uc.repo.CreateFavorite(userID, movieID)
	if err != nil {
		return false, dtos.NewRequestError("Error while favoriting movie")
	}
	return true, nil // true = foi adicionado
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

