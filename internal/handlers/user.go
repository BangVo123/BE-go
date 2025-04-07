package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetMe(c *gin.Context)
	GetUserInfo(c *gin.Context)
}
