package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port                     string `env:"PORT"`
	ClientUrl                string `env:"CLIENT_URL"`
	SessionSecret            string `env:"SESSION_SECRET"`
	JWTAccessTokenSecret     string `env:"JWT_ACCESS_TOKEN_SECRET"`
	JWTRefreshTokenSecret    string `env:"JWT_REFRESH_TOKEN_SECRET"`
	JWTAccessTokenExpiry     string `env:"JWT_ACCESS_TOKEN_EXPIRY"`
	JWTRefreshTokenExpiry    string `env:"JWT_REFRESH_TOKEN_EXPIRY"`
	Email                    string `env:"EMAIL"`
	EmailPassword            string `env:"EMAIL_PASSWORD"`
	CloudinaryCloudName      string `env:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryCloudAPIKey    string `env:"CLOUDINARY_CLOUD_API_KEY"`
	CloudinaryCloudAPISecret string `env:"CLOUDINARY_CLOUD_API_SECRET"`
	GoogleClientID           string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret       string `env:"GOOGLE_CLIENT_SECRET"`
	GoogleCallbackURL        string `env:"GOOGLE_CALLBACK_URL"`
	ScopeURL                 string `env:"SCOPE_URL"`
	MongoDbURL               string `env:"MONGODB_URL"`
	MongoDbName              string `env:"MONGODB_NAME"`
}

func NewConfig() *Configuration {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}

	cfg := Configuration{}

	err = env.Parse(&cfg)

	if err != nil {
		log.Fatal("Unable to parse environment variable: ", err)
	}

	return &cfg
}
