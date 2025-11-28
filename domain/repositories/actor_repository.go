package repositories

import (
	"cinelist/domain/entities"
)

type ActorRepository interface {
	GetAll() ([]entities.Actor, error)
}


