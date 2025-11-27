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

type MovieDetailData struct {
	Movie   entities.Movie     `json:"movie"`
	Cast    []entities.Actor   `json:"cast"`
	Ratings []entities.Watched `json:"ratings"`
}

type MovieDetailResponseDTO struct {
	Success bool            `json:"success"`
	Data    MovieDetailData `json:"data"`
}
