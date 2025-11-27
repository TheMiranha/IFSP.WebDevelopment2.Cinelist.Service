package repositories

import (
	"cinelist/domain/entities"
)

type MovieRepository interface {
	GetAll() ([]entities.Movie, error)
	SearchByTitle(term string) ([]entities.Movie, error)
	GetCastByMovieID(movieID string) ([]entities.Actor, error)
	GetClassics() ([]entities.Movie, error)
	GetHighlights() ([]entities.Movie, error)
	GetNewReleases() ([]entities.Movie, error)
	GetById(id string) (entities.Movie, error)
	GetRatingsWithUserByMovieID(movieID string) ([]RatingWithUser, error)
}

type RatingWithUser struct {
	UserID      string
	UserName    string
	UserImageUrl string
	Rate        float64
	Description string
	CreatedAt   string
	UpdatedAt   string
}

