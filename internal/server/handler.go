package server

import (
	"project/internal/delivery/http"
	"project/internal/middlewares"
	"project/internal/repositories"
	"project/internal/services"
	"project/routers"

	"github.com/gin-gonic/gin"
)

func (s *Server) Handler(g *gin.Engine, mongoStore repositories.MongoSessionStore) {
	db := s.db.Database(s.cfg.MongoDbName)

	AuthRepo := repositories.NewUserRepo(db, "users")
	VideoRepo := repositories.NewVideoRepo(db, "videos")

	AuthService := services.NewAuthService(AuthRepo)
	VideoService := services.NewVideoService(VideoRepo)

	AuthHandler := http.NewAuthHandler(AuthService, s.cfg, &mongoStore)
	UserHandler := http.NewAuthHandler(AuthService, s.cfg, &mongoStore)
	VideoHandler := http.NewVideoHandler(VideoService)

	mw := middlewares.NewMiddlewareManager(AuthService, s.cfg, s.MongoStore)

	v1 := g.Group("/api/v1")
	authGroup := v1.Group("/auth")
	userGroup := v1.Group("/users")
	videoGroup := v1.Group("/video")

	routers.MapAuthRoute(authGroup, AuthHandler, mw)
	routers.MapUserRoute(userGroup, UserHandler, mw)
	routers.MapVideoRoute(videoGroup, VideoHandler, mw)

}
