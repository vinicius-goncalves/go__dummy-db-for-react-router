package methods

import (
	"api/database"
	"api/structs"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() []byte {

	client := database.Connect()

	col := client.Database("users").Collection("users")

	cursor, err := col.Find(context.Background(), bson.D{})

	if err != nil {
		panic(err)
	}

	users := new([]structs.User)

	if err = cursor.All(context.TODO(), users); err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		panic(err)
	}

	return j
}
