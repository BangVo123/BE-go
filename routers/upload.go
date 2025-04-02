package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func UploadRoute(videoGroup *gin.RouterGroup, h handlers.UploadHandler, mw *middlewares.MiddlewareManager) {
	videoGroup.POST("", mw.Protect, h.VideoUpload)
	videoGroup.POST("/avatar", mw.Protect, h.AvatarUpload)
}
