package main

import (
	"api/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.StartServer()

}
