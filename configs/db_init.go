package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB(ctx *context.Context) (*mongo.Client, error) {
	err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("MONGODB_NAME")
	host := os.Getenv("MONGODB_HOST")
	password := os.Getenv("MONGODB_PASSWOARD")

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s", user, password, host))
	client, err := mongo.Connect(*ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if client.Ping(*ctx, readpref.Primary()) != nil {
		return nil, err
	}
	return client, nil
}
