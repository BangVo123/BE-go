package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func MapAuthRoute(authGroup *gin.RouterGroup, h handlers.AuthHandler, mw *middlewares.MiddlewareManager) {
	authGroup.POST("/login", h.Login)
	authGroup.GET("/:provider", h.GoogleOauth)
	authGroup.GET("/:provider/callback", h.GoogleOauthCallback)
	authGroup.GET("/logout", mw.Protect, h.Logout)
}
