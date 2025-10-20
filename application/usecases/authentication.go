package usecases

import (
	"cinelist/domain/dtos"
	"cinelist/domain/entities"
	infrastructure_utils "cinelist/infrastructure"
	"cinelist/infrastructure/database/repositories"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AuthenticationUseCase struct {
	repo repositories.UserRepository
}

func NewAuthenticationUseCase(repo repositories.UserRepository) AuthenticationUseCase {
	return AuthenticationUseCase{repo: repo}
}

func (u *AuthenticationUseCase) SignIn(payload dtos.SignInDTO) (dtos.SignInResponseDTO, *dtos.RequestError) {
	user, err := u.repo.GetByEmail(payload.Email)

	if err != nil {
		return dtos.SignInResponseDTO{}, dtos.NewRequestError("Invalid credentials")
	}

	samePassword := infrastructure_utils.CheckPasswordHash(payload.Password, user.Password)

	if !samePassword {
		return dtos.SignInResponseDTO{}, dtos.NewRequestError("Invalid credentials")
	}

	token, err := infrastructure_utils.GenerateJWT(user.ID)

	if err != nil {
		return dtos.SignInResponseDTO{}, dtos.NewRequestError("Error on token generation")
	}

	return dtos.SignInResponseDTO{
		Success:     true,
		AccessToken: token,
	}, nil
}

func (u *AuthenticationUseCase) Create(payload dtos.SignUpDTO) (dtos.SignUpResponseDTO, *dtos.RequestError) {

	_, alreadyExists := u.repo.GetByEmail(payload.Email)

	if alreadyExists == nil {
		return dtos.SignUpResponseDTO{}, dtos.NewRequestError("Email already in use")
	}

	hashedPassword, err := infrastructure_utils.HashPassword(payload.Password)

	if err != nil {
		return dtos.SignUpResponseDTO{}, &dtos.RequestError{Success: false, Message: "Error on password hashing"}
	}

	user := entities.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Password:  hashedPassword,
		Email:     payload.Email,
		ImageUrl:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	u.repo.Create(user)

	token, err := infrastructure_utils.GenerateJWT(user.ID)

	if err != nil {
		fmt.Println(err)
		return dtos.SignUpResponseDTO{}, &dtos.RequestError{Success: false, Message: "Error on token generation"}
	}

	return dtos.SignUpResponseDTO{
		Success:     true,
		AccessToken: token,
	}, nil
}
