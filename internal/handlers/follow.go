package handlers

import (
	"github.com/gin-gonic/gin"
)

type FollowHandler interface {
	Follow(c *gin.Context)
}
