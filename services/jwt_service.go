package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/RPA_VoucherExchange/entities"
	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	EmployeeID uint `json:"employee_id"`
	ProviderID uint `json:"provider_id"`
	jwt.RegisteredClaims
}

func GenerateToken(employee entities.Employee) (string, error) {
	claims := jwtCustomClaims{
		EmployeeID: employee.Model.ID,
		ProviderID: employee.ProviderID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), //1 week
			Issuer:    "github.com/ntikhoa",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte("github.com/RPA_VoucherExchange/superlongsecretkey"))
	if err != nil {
		return "", errors.New("cannot create token")
	}
	return token, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("github.com/RPA_VoucherExchange/superlongsecretkey"), nil
	})
}
