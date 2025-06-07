package config

import (
	"os"
	"path/filepath"
	"recipebook/core/logic"
	"strconv"

	"github.com/joho/godotenv"
)

var DataDirectory string
var ImagesDirectory string

func LoadEnv() {
	godotenv.Load(".env")

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
