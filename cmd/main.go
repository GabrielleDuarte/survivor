package main

import (
	"context"
	"dataapi/internal/config"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	//Pega a variável de ambiente MONGO_URI
	mongoURI := cfg.Mongo.URI

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	databaseName := os.Getenv("MONGO_DATABASE_NAME")

	fmt.Printf("✅ Conectado ao MongoDB (%s)\n", databaseName)
}
