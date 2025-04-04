package handlers

import "github.com/gin-gonic/gin"

type CodeHandler interface {
	GetCode(c *gin.Context)
	CreateCode(c *gin.Context)
	DeleteCode(c *gin.Context)
}
