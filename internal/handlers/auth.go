package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	GoogleOauth(c *gin.Context)
	GoogleOauthCallback(c *gin.Context)
	Logout(c *gin.Context)
	GetMe(c *gin.Context)
	GetUserInfo(c *gin.Context)
}
