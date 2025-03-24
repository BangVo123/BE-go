package utils

import (
	"fmt"
	"project/internal/presenter"

	"github.com/golang-jwt/jwt"
)

func GenToken(payload map[string]string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    payload["_id"],
		"email": payload["email"],
		// "exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, secret string) (*presenter.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &presenter.TokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return &presenter.TokenClaims{}, err
	}

	if claims, ok := token.Claims.(*presenter.TokenClaims); !ok || !token.Valid {
		return &presenter.TokenClaims{}, fmt.Errorf("Invalid token")
	} else {
		return claims, nil
	}
}
