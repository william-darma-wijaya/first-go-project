package helper

import (
	"first-go-project/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(authConfig *entity.AuthConfig, userID string) (string, string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(authConfig.MinutesExp)).Unix(),
		"type": "access",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccess, err := token.SignedString([]byte(authConfig.Secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Minute * time.Duration(authConfig.RefreshMinutesExp)).Unix(),
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefresh, err := refreshToken.SignedString([]byte(authConfig.Secret))
	if err != nil {
		return "", "", err
	}

	return signedAccess, signedRefresh, nil
}
