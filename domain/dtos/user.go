package dtos

import "cinelist/domain/entities"

type UserData struct {
	User      entities.User      `json:"user"`
	Favorites []entities.Movie   `json:"favorites"`
	ToWatch   []entities.Movie   `json:"toWatch"`
	Watched   []entities.Watched `json:"watched"`
}

type UserResponseDTO struct {
	Success bool     `json:"success"`
	Data    UserData `json:"data"`
}

