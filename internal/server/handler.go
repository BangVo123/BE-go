package server

import (
	"project/internal/delivery/http"
	"project/internal/middlewares"
	"project/internal/repositories"
	"project/internal/services"
	"project/routers"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	mail "gopkg.in/gomail.v2"
)

func (s *Server) Handler(g *gin.Engine, mongoStore repositories.MongoSessionStore, cld cloudinary.Cloudinary, dialer *mail.Dialer) {
	db := s.db.Database(s.cfg.MongoDbName)

	AuthRepo := repositories.NewUserRepo(db, "users")
	VideoRepo := repositories.NewVideoRepo(db, "videos")
	VideoWithOwnerInfoRepo := repositories.NewVideoWithOwnerInfoRepo(db, "videos")
	CodeRepo := repositories.NewCodeRepo(db, "digitcodes")
	FollowRepo := repositories.NewFollowerRepo(db, "followers")
	FollowerRepo := repositories.NewFollowerInfoRepo(db, "followers")
	FollowingRepo := repositories.NewFollowingInfoRepo(db, "followers")
	LikeRepo := repositories.NewLikeRepo(db, "likes")
	LoveRepo := repositories.NewLoveRepo(db, "loves")

	AuthService := services.NewAuthService(AuthRepo)
	VideoService := services.NewVideoService(VideoRepo)
	VideoWithOwnerInfoService := services.NewVideoWithOwnerInfoService(VideoWithOwnerInfoRepo)
	CodeService := services.NewCodeService(CodeRepo)
	FollowService := services.NewFollowService(FollowRepo)
	FollowerService := services.NewFollowerInfoService(FollowerRepo)
	FollowingService := services.NewFollowingInfoService(FollowingRepo)
	LikeService := services.NewLikeService(LikeRepo)
	LoveService := services.NewLoveService(LoveRepo)

	AuthHandler := http.NewAuthHandler(AuthService, s.cfg, &mongoStore, CodeService, *dialer)
	UserHandler := http.NewUserHandler(AuthService, FollowService, FollowerService, FollowingService, LikeService, LoveService)
	VideoHandler := http.NewVideoHandler(VideoService, cld, LikeService, LoveService)
	VideoWithOwnerInfoHandler := http.NewVideoWithOwnerInfoHandler(VideoWithOwnerInfoService)
	UploadHandler := http.NewUploadHandler(cld)
	FollowHandler := http.NewFollowhandler(FollowService)

	mw := middlewares.NewMiddlewareManager(AuthService, s.cfg, s.MongoStore)

	v1 := g.Group("/api/v1")
	authGroup := v1.Group("/auth")
	userGroup := v1.Group("/users")
	videoGroup := v1.Group("/video")
	uploadGroup := v1.Group("/upload")
	followGroup := v1.Group("/follow")

	routers.MapAuthRoute(authGroup, AuthHandler, mw)
	routers.MapUserRoute(userGroup, UserHandler, mw)
	routers.MapVideoWithOwnerInfoRoute(videoGroup, VideoWithOwnerInfoHandler, mw)
	routers.MapVideoRoute(videoGroup, VideoHandler, mw)
	routers.UploadRoute(uploadGroup, UploadHandler, mw)
	routers.FollowRoute(followGroup, FollowHandler, mw)

}
