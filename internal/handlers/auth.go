package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login() gin.HandlerFunc
	SignUp() gin.HandlerFunc
	GoogleOauth() gin.HandlerFunc
	GoogleOauthCallback() gin.HandlerFunc
}
