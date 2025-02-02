package logic

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

var DataDirectory string
var ImagesDirectory string

func LoadEnv() {
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "./data"
	}

	DataDirectory, _ = filepath.Abs(dataPath)
	CreateDir(DataDirectory)

	ImagesDirectory, _ = filepath.Abs(filepath.Join(DataDirectory, "images"))
	CreateDir(ImagesDirectory)
}

func IsLocalDevEnv() bool {
	localDev := os.Getenv("LOCAL_DEV_ENV")
	localDevBool, _ := strconv.ParseBool(localDev)
	return localDevBool
}

func CreateDir(directoryPath string) {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		log.Printf("Creating directory: %s", directoryPath)
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			log.Printf("Error creating directory%s: %s", directoryPath, err)
		} else {
			log.Printf("Directory created: %s", directoryPath)
		}
	} else {
		log.Printf("Directory already exists: %s", directoryPath)
	}
}

func StringArraySortUnique(arrayToSort []string) []string {
	slices.Sort(arrayToSort)
	arrayToSort = slices.Compact(arrayToSort)
	return arrayToSort
}
