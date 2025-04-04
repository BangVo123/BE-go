package routers

import (
	"project/internal/handlers"
	"project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func MapAuthRoute(authGroup *gin.RouterGroup, h handlers.AuthHandler, mw *middlewares.MiddlewareManager) {
	authGroup.POST("/login", h.Login)
	authGroup.POST("/signup", h.SignUp)
	authGroup.GET("/:provider", h.GoogleOauth)
	authGroup.GET("/:provider/callback", h.GoogleOauthCallback)
	authGroup.GET("/logout", mw.Protect, h.Logout)
	authGroup.POST("/verify", h.Forgot)
	authGroup.POST("/reset", h.Reset)
}
