package structs

type User struct {
	Id             string `bson:"_id"`
	Name           string `bson:"name,omitempty"`
	GamesOwned     int    `bson:"games_owned"`
	MostPlayedGame string `bson:"most_played_game"`
}
