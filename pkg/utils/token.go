package utils

import (
	"time"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User, secretKey string, ttl int) (string, error) {
	jwtSecret := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"role_id":  user.RoleID,
		"exp":      time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
