package main

import (
	"recipebook/core/config"
	"recipebook/core/database"
	"recipebook/core/logic"
)

func main() {
	logic.GetBootTime()
	config.LoadEnv()
	database.Initialise()
	StartServer()
}
