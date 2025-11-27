package repositories

import (
	"cinelist/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetById(id uuid.UUID) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
}

