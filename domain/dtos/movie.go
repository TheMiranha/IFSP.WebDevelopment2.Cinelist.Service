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