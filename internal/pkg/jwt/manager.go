package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"shop-smart-api/internal/entity"
	"time"
)

type Manager interface {
	Generate(user *entity.User, isFully bool, role string) (string, error)
	Verify(accessToken string) (*UserClaims, error)
}

type jwtManager struct {
	secret string
}

func CreateManager(secret string) Manager {
	return &jwtManager{secret}
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserRole string `json:"role"`
	UserId   string `json:"user_id"`
	IsFully  bool   `json:"is_fully"`
}

func (j *jwtManager) Generate(user *entity.User, isFully bool, role string) (string, error) {

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		UserRole: role,
		UserId:   user.ID,
		IsFully:  isFully,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secret))
}

func (j *jwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}

		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
