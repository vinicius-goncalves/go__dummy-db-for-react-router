package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CLIENT_URI string = os.Getenv("CLIENT_URI")

func Connect() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(CLIENT_URI))

	if err != nil {
		panic(err)
	}

	return client
}
