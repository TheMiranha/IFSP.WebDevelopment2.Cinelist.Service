package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	"cinelist/infrastructure/database/repositories"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUserById(id uuid.UUID) (entities.User, *dtos.RequestError) {
	user, err := uc.repo.GetById(id)

	if err != nil {
		return entities.User{}, dtos.NewRequestError("User not found")
	}

	return user, nil
}
