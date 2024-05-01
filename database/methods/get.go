package methods

import (
	"api/database"
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(id string) []byte {

	client := database.Connect()

	col := client.Database("users").Collection("users")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	var r bson.M

	if err = col.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&r); err != nil {

		if err == mongo.ErrNoDocuments {
			fmt.Print("No documents were found.")
			return nil
		}

		panic(err)
	}

	j, err := json.MarshalIndent(r, "", " ")

	if err != nil {
		panic(err)
	}

	return j
}
