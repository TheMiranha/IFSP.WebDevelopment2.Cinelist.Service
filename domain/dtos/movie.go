package dtos

import "cinelist/domain/entities"

type MoviesResponseDTO struct {
	Success bool             `json:"success"`
	Data    []entities.Movie `json:"data"`
}

type MovieWithCast struct {
	Movie entities.Movie   `json:"movie"`
	Cast  []entities.Actor `json:"cast"`
}

type MoviesWithCastResponseDTO struct {
	Success bool            `json:"success"`
	Data    []MovieWithCast `json:"data"`
}

type TopMoviesData struct {
	Classics    []entities.Movie `json:"classics"`
	Highlights  []entities.Movie `json:"highlights"`
	NewReleases []entities.Movie `json:"newReleases"`
}

type TopMoviesResponseDTO struct {
	Success bool          `json:"success"`
	Data    TopMoviesData `json:"data"`
}

type RatingDTO struct {
	User        RatingUserDTO `json:"user"`
	Rate        float64       `json:"rate"`
	Description string        `json:"description"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
}

type RatingUserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

type MovieDetailData struct {
	Movie   entities.Movie `json:"movie"`
	Cast    []entities.Actor `json:"cast"`
	Ratings []RatingDTO    `json:"ratings"`
}

type MovieDetailResponseDTO struct {
	Success bool            `json:"success"`
	Data    MovieDetailData `json:"data"`
}

type FavoriteMovieDTO struct {
	MovieId string `json:"movieId" binding:"required"`
}

type ToWatchMovieDTO struct {
	MovieId string `json:"movieId" binding:"required"`
}

type WatchedMovieDTO struct {
	MovieId    string  `json:"movieId" binding:"required"`
	Rate       float64 `json:"rate" binding:"required"`
	Description string `json:"description"`
}

type SuccessResponseDTO struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
