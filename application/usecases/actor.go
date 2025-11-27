package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	domain_repositories "cinelist/domain/repositories"
)

type ActorUseCase struct {
	repo domain_repositories.ActorRepository
}

func NewActorUseCase(repo domain_repositories.ActorRepository) ActorUseCase {
	return ActorUseCase{repo: repo}
}

func (uc *ActorUseCase) GetAll() ([]entities.Actor, *dtos.RequestError) {
	actors, err := uc.repo.GetAll()

	if err != nil {
		return []entities.Actor{}, dtos.NewRequestError("Error while selecting all actors")
	}

	return actors, nil
}
