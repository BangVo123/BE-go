package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func FollowRoute(followGroup *gin.RouterGroup, h handlers.FollowHandler, mw *middlewares.MiddlewareManager) {
	followGroup.POST("", mw.Protect, h.Follow)
}
