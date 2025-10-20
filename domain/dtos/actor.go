package dtos

import "cinelist/domain/entities"

type ActorsResponseDTO struct {
	Success bool             `json:"success"`
	Data    []entities.Actor `json:"data"`
}
