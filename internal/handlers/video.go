package handlers

import "github.com/gin-gonic/gin"

type VideoHandler interface {
	GetVideos(c *gin.Context)
	// GetVideos() gin.HandlerFunc
}
