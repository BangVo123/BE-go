package server

import (
	"project/config"
	"project/internal/delivery/http"
	"project/internal/middlewares"
	"project/internal/repositories"
	"project/internal/services"
	"project/routers"

	"github.com/gin-gonic/gin"
)

func (s *Server) Handler(g *gin.Engine, mongoStore repositories.MongoSessionStore) {
	db := s.db.Database(s.cfg.MongoDbName)
	cfg := config.NewConfig()

	AuthRepo := repositories.NewUserRepo(db, "users")

	AuthService := services.NewAuthService(*AuthRepo)

	AuthHandler := http.NewAuthHandler(AuthService, cfg, &mongoStore)
	mw := middlewares.NewMiddlewareManager(AuthService, s.cfg, s.MongoStore)

	v1 := g.Group("/api/v1")
	authGroup := v1.Group("/auth")

	routers.MapAuthRoute(authGroup, AuthHandler, mw)

}
