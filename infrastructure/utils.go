package infrastructure_utils

import (
	"cinelist/domain/dtos"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type JwtClaims struct {
	UserID uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uuid.UUID) (string, error) {
	godotenv.Load()

	claims := &JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
	}

	var secretKey []byte = []byte(os.Getenv("JWT_SIGNATURE"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ThrowInvalidRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, dtos.NewRequestError("Invalid request"))
}
