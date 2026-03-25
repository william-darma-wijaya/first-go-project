package helper

import (
	"first-go-project/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(authConfig *entity.AuthConfig, userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(authConfig.MinutesExp)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(authConfig.Secret))
}
