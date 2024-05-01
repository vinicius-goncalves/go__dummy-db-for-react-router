package structs

type EditUser struct {
	Id             string `json:"id"`
	GamesOwned     int    `json:"games_owned,string"`
	MostPlayedGame string `json:"most_played_game"`
}
