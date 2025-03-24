package routers

import (
	auth "project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func MapAuthRoute(authGroup *gin.RouterGroup, h auth.Handler) {
	authGroup.POST("/login", h.Login())
	authGroup.GET("/:provider", h.GoogleOauth())
	authGroup.GET("/:provider/callback", h.GoogleOauthCallback())
}
