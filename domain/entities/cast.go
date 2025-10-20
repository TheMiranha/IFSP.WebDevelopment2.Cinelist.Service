package entities

import (
	"time"

	"github.com/google/uuid"
)

type Cast struct {
	Actor     uuid.UUID `json:"actorId"`
	Movie     uuid.UUID `json:"movieId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
