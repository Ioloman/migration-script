package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var CTX = context.TODO()

func SetupDB() {
	log.Println("Connecting to MongoDB")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(CTX, nil)
	if err != nil {
		panic(err)
	}

	Collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_COLLECTION"))
	log.Println("Connected to MongoDB")
}
