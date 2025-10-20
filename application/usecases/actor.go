package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	"cinelist/infrastructure/database/repositories"
)

type ActorUseCase struct {
	repo repositories.ActorRepository
}

func NewActorUseCase(repo repositories.ActorRepository) ActorUseCase {
	return ActorUseCase{repo: repo}
}

func (uc *ActorUseCase) GetAll() ([]entities.Actor, *dtos.RequestError) {
	actors, err := uc.repo.GetAll()

	if err != nil {
		return []entities.Actor{}, dtos.NewRequestError("Error while selecting all actors")
	}

	return actors, nil
}
