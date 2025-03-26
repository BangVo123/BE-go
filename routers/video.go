package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func MapVideoRoute(videoGroup *gin.RouterGroup, h handlers.VideoHandler, mw *middlewares.MiddlewareManager) {
	videoGroup.GET("", h.GetVideos)
}
