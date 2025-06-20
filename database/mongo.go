package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ProductCollection *mongo.Collection

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB not responding:", err)
	}

	dbName := os.Getenv("DB_NAME")
	ProductCollection = client.Database(dbName).Collection("products")
	log.Println("✅ Connected to MongoDB!")
}