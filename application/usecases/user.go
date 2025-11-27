package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	"cinelist/infrastructure/database/repositories"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo              repositories.UserRepository
	interactionRepo   repositories.MovieInteractionRepository
}

func NewUserUseCase(repo repositories.UserRepository, interactionRepo repositories.MovieInteractionRepository) UserUseCase {
	return UserUseCase{
		repo:            repo,
		interactionRepo: interactionRepo,
	}
}

func (uc *UserUseCase) GetUserById(id uuid.UUID) (dtos.UserData, *dtos.RequestError) {
	user, err := uc.repo.GetById(id)

	if err != nil {
		return dtos.UserData{}, dtos.NewRequestError("User not found")
	}

	favorites, err := uc.interactionRepo.GetFavoritesByUserID(id)
	if err != nil {
		favorites = []entities.Movie{}
	}

	toWatch, err := uc.interactionRepo.GetToWatchByUserID(id)
	if err != nil {
		toWatch = []entities.Movie{}
	}

	watched, err := uc.interactionRepo.GetWatchedByUserID(id)
	if err != nil {
		watched = []entities.Watched{}
	}

	return dtos.UserData{
		User:      user,
		Favorites: favorites,
		ToWatch:   toWatch,
		Watched:   watched,
	}, nil
}
