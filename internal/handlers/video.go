package handlers

import "github.com/gin-gonic/gin"

type VideoHandler interface {
	GetVideos(c *gin.Context)
	AddVideo(c *gin.Context)
	Like(c *gin.Context)
	Love(c *gin.Context)
}

type VideoWithOwnerInfoHandler interface {
	GetVideos(c *gin.Context)
}
