package entities

import (
	"time"

	"github.com/google/uuid"
)

type ToWatch struct {
	User      uuid.UUID `json:"userId"`
	Movie     uuid.UUID `json:"movieId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
