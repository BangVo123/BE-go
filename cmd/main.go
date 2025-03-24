package main

import (
	"context"
	"log"
	"project/config"
	"project/internal/server"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.NewConfig()

	log.Println("url: ", cfg.MongoDbURL)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, error := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDbURL))
	if error != nil {
		log.Fatal("MongoDB connection error: ", error)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB pinging error: ", err)
	}

	log.Println("MongoDB connection success")

	srv := server.New(cfg, client)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
