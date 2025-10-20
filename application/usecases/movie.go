package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	"cinelist/infrastructure/database/repositories"
)

type MovieUseCase struct {
	repo repositories.MovieRepository
}

func NewMovieUseCase(repo repositories.MovieRepository) MovieUseCase {
	return MovieUseCase{repo: repo}
}

func (uc *MovieUseCase) GetAll() ([]entities.Movie, *dtos.RequestError) {
	movies, err := uc.repo.GetAll()

	if err != nil {
		return []entities.Movie{}, dtos.NewRequestError("Error while selecting all movies")
	}

	return movies, nil
}
