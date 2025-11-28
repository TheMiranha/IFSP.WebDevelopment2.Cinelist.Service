package services

import (
	"github.com/google/uuid"
)

type AuthService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	GenerateJWT(userID uuid.UUID) (string, error)
}


