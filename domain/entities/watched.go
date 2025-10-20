package entities

import (
	"time"

	"github.com/google/uuid"
)

type Watched struct {
	User        uuid.UUID `json:"userId"`
	Movie       uuid.UUID `json:"movieId"`
	Rate        float64   `json:"rate"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
