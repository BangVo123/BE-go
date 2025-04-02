package main

import (
	"context"
	"log"
	"project/config"
	"project/internal/repositories"
	"project/internal/server"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.NewConfig()

	log.Println("url: ", cfg.MongoDbURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, error := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDbURL))
	if error != nil {
		log.Fatal("MongoDB connection error: ", error)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB pinging error: ", err)
	}

	log.Println("MongoDB connection success")

	//create mongo store here
	mongoStore := repositories.NewMongoSessionStore(client, cfg.MongoDbName, "sessions")

	//create connection to cloudinary
	cld, error := cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryCloudAPIKey, cfg.CloudinaryCloudAPISecret)
	if error != nil {
		log.Fatal("Cloudinary connection error: ", error)
	}

	srv := server.New(cfg, client, mongoStore, cld)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
