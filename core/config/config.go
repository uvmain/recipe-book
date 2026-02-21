package config

import (
	"log"
	"os"
	"path/filepath"
	"recipebook/core/logic"
	"strconv"

	"github.com/joho/godotenv"
)

var DataDirectory string
var ImagesDirectory string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("No .env file found, relying on environment variables")
	}

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "./data"
	}

	DataDirectory, _ = filepath.Abs(dataPath)
	logic.CreateDir(DataDirectory)

	ImagesDirectory, _ = filepath.Abs(filepath.Join(DataDirectory, "images"))
	logic.CreateDir(ImagesDirectory)
}

func IsLocalDevEnv() bool {
	localDev := os.Getenv("LOCAL_DEV_ENV")
	localDevBool, _ := strconv.ParseBool(localDev)
	return localDevBool
}
