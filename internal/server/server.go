package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/config"
	"project/internal/repositories"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	gin        *gin.Engine
	cfg        *config.Configuration
	db         *mongo.Client
	MongoStore *repositories.MongoSessionStore
}

func New(cfg *config.Configuration, db *mongo.Client, MongoStore *repositories.MongoSessionStore) *Server {
	return &Server{gin: gin.New(), cfg: cfg, db: db, MongoStore: MongoStore}
}

func (s *Server) Run() error {
	s.gin.OPTIONS("/*cors", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{s.cfg.ClientUrl}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	s.gin.Use(cors.New(config))

	srv := http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: s.gin.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Something went wrong when started server. ", err)
		}
	}()

	//create google provider for google oauth
	goth.UseProviders(
		google.New(s.cfg.GoogleClientID, s.cfg.GoogleClientSecret, s.cfg.GoogleCallbackURL, "email", "profile"),
	)

	s.Handler(s.gin, *s.MongoStore)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown")
	}

	<-ctx.Done()
	log.Println("Server exiting")

	return nil
}
