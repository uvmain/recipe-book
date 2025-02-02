package main

import (
	"recipebook/database"
	"recipebook/logic"
)

func main() {
	logic.LoadEnv()
	database.Initialise()
	StartServer()
}
