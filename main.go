package main

import (
	"recipebook/core/config"
	"recipebook/core/database"
)

func main() {
	config.LoadEnv()
	database.Initialise()
	StartServer()
}
