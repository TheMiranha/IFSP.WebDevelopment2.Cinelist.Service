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

func (uc *MovieUseCase) GetTopMovies() (dtos.TopMoviesData, *dtos.RequestError) {
	classics, err := uc.repo.GetClassics()
	if err != nil {
		return dtos.TopMoviesData{}, dtos.NewRequestError("Error while fetching classics")
	}

	highlights, err := uc.repo.GetHighlights()
	if err != nil {
		return dtos.TopMoviesData{}, dtos.NewRequestError("Error while fetching highlights")
	}

	newReleases, err := uc.repo.GetNewReleases()
	if err != nil {
		return dtos.TopMoviesData{}, dtos.NewRequestError("Error while fetching new releases")
	}

	return dtos.TopMoviesData{
		Classics:    classics,
		Highlights:  highlights,
		NewReleases: newReleases,
	}, nil
}

func (uc *MovieUseCase) GetMovieById(id string) (dtos.MovieDetailData, *dtos.RequestError) {
	movie, err := uc.repo.GetById(id)
	if err != nil {
		return dtos.MovieDetailData{}, dtos.NewRequestError("Movie not found")
	}

	cast, err := uc.repo.GetCastByMovieID(id)
	if err != nil {
		cast = []entities.Actor{}
	}

	ratings, err := uc.repo.GetWatchedByMovieID(id)
	if err != nil {
		ratings = []entities.Watched{}
	}

	return dtos.MovieDetailData{
		Movie:   movie,
		Cast:    cast,
		Ratings: ratings,
	}, nil
}
