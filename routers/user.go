package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func MapUserRoute(authGroup *gin.RouterGroup, h handlers.AuthHandler, mw *middlewares.MiddlewareManager) {
	authGroup.GET("/:userId", h.GetUserInfo)
	authGroup.Use(mw.Protect)
	authGroup.GET("", h.GetMe)

}
