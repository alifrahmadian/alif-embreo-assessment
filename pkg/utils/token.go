package utils

import (
	"time"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	RoleID    int64  `json:"role_id"`
	CompanyID *int64 `json:"company_id,omitempty"`
	VendorID  *int64 `json:"vendor_id,omitempty"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, secretKey string, ttl int) (string, error) {
	jwtSecret := []byte(secretKey)

	claims := Claims{
		Username:  user.Username,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CompanyID: user.CompanyID,
		VendorID:  user.VendorID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(ttl))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
