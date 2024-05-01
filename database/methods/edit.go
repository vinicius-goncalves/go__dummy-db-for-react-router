package methods

import (
	"api/database"
	"api/structs"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Edit(id string, body structs.EditUser) {

	client := database.Connect()
	col := client.Database("users").Collection("users")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "games_owned", Value: body.GamesOwned},
		{Key: "most_played_game", Value: body.MostPlayedGame},
	}}}

	if _, err = col.UpdateOne(context.TODO(), filter, update); err != nil {
		panic(err)
	}
}
