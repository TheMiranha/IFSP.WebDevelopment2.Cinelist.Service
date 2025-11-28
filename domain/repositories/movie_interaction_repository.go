package repositories

import (
	"cinelist/domain/entities"

	"github.com/google/uuid"
)

type MovieInteractionRepository interface {
	CreateFavorite(userID uuid.UUID, movieID uuid.UUID) error
	DeleteFavorite(userID uuid.UUID, movieID uuid.UUID) error
	GetFavoriteByUserAndMovie(userID uuid.UUID, movieID uuid.UUID) (entities.Favorite, error)
	CreateToWatch(userID uuid.UUID, movieID uuid.UUID) error
	DeleteToWatch(userID uuid.UUID, movieID uuid.UUID) error
	CreateWatched(watched entities.Watched) error
	UpdateWatched(watched entities.Watched) error
	GetWatchedByUserAndMovie(userID uuid.UUID, movieID uuid.UUID) (entities.Watched, error)
	GetFavoritesByUserID(userID uuid.UUID) ([]entities.Movie, error)
	GetToWatchByUserID(userID uuid.UUID) ([]entities.Movie, error)
	GetWatchedByUserID(userID uuid.UUID) ([]entities.Watched, error)
}


