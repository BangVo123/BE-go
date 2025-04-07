package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func MapVideoWithOwnerInfoRoute(videoGroup *gin.RouterGroup, h handlers.VideoWithOwnerInfoHandler, mw *middlewares.MiddlewareManager) {
	videoGroup.GET("", h.GetVideos)
}

func MapVideoRoute(videoGroup *gin.RouterGroup, h handlers.VideoHandler, mw *middlewares.MiddlewareManager) {
	videoGroup.POST("", mw.Protect, h.AddVideo)
	videoGroup.POST("/like/:videoId", mw.Protect, h.Like)
	videoGroup.POST("/love/:videoId", mw.Protect, h.Love)
}
