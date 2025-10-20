package dtos

import "cinelist/domain/entities"

type MoviesResponseDTO struct {
	Success bool             `json:"success"`
	Data    []entities.Movie `json:"data"`
}
