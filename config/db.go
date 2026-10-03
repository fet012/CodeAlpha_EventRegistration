package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnect()*mongo.Database{
	godotenv.Load()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if (err){
		log.Fatal("Database connection has failed", err)
	}
	return (
		client.Database(os.Getenv("DB_NAME"))
	)
}