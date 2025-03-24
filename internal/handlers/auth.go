package auth

import "github.com/gin-gonic/gin"

type Handler interface {
	Login() gin.HandlerFunc
	SignUp() gin.HandlerFunc
	GoogleOauth() gin.HandlerFunc
	GoogleOauthCallback() gin.HandlerFunc
}
