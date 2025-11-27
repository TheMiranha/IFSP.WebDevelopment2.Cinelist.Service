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

func (uc *MovieUseCase) GetAllByTitle(term string) ([]entities.Movie, *dtos.RequestError) {
	movies, err := uc.repo.SearchByTitle(term)

	if err != nil {
		return []entities.Movie{}, dtos.NewRequestError("Error while searching movies")
	}

	return movies, nil
}

func (uc *MovieUseCase) Search(term string) ([]dtos.MovieWithCast, *dtos.RequestError) {
	movies, err := uc.repo.SearchByTitle(term)

	if err != nil {
		return []dtos.MovieWithCast{}, dtos.NewRequestError("Error while searching movies")
	}

	moviesWithCast := make([]dtos.MovieWithCast, 0)

	for _, movie := range movies {
		cast, err := uc.repo.GetCastByMovieID(movie.ID.String())
		if err != nil {
			cast = []entities.Actor{}
		}

		moviesWithCast = append(moviesWithCast, dtos.MovieWithCast{
			Movie: movie,
			Cast:  cast,
		})
	}

	return moviesWithCast, nil
}