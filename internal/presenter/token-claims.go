package presenter

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	ID    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"email"`
	jwt.StandardClaims
}
