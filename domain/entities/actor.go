package entities

import (
	"time"

	"github.com/google/uuid"
)

type Actor struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
