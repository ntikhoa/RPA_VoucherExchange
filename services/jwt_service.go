package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(account entities.Account) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

//convert uint ID into string
//because jwt library convert uint to float64
//which cause trouble to manipulate data
type jwtCustomClaims struct {
	AccountID  string `json:"account_id"`
	ProviderID string `json:"provider_id"`
	IssueAt    string `json:"issue_at"`
	jwt.RegisteredClaims
}

func (s *jwtService) GenerateToken(account entities.Account) (string, error) {

	claims := jwtCustomClaims{
		AccountID:  strconv.FormatUint(uint64(account.Model.ID), 10),
		ProviderID: strconv.FormatUint(uint64(account.ProviderID), 10),
		IssueAt:    account.IssueAt.Format(constants.JWT_DATE_FORMAT),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), //1 week
			Issuer:    "github.com/ntikhoa",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(os.Getenv("JWT_SIGNED_KEY")))
	if err != nil {
		log.Println(err)
		return "", errors.New("cannot create token")
	}
	return token, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SIGNED_KEY")), nil
	})
}
