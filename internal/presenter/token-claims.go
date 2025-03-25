package presenter

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	Id      string `json:"id" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Expired int64  `json:"exp"`
	jwt.StandardClaims
}
