package logic

import (
	"errors"
	"log"
	"os"
	"recipebook/core/types"
	"slices"
	"strconv"
	"strings"
)

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

func ConvertNullStringToInt(nullStringValue types.NullString) (int, error) {
	if !nullStringValue.Valid {
		return 0, nil
	}
	integerValue, err := strconv.Atoi(nullStringValue.String)
	if err != nil {
		return 0, errors.New("failed to convert string to int")
	}
	return integerValue, nil
}

func FilterEmptyStringsFromSlice(stringSlice []string) []string {
	if len(stringSlice) > 0 {
		nonEmptySlice := make([]string, 0, len(stringSlice))
		for _, course := range stringSlice {
			if strings.TrimSpace(course) != "" {
				nonEmptySlice = append(nonEmptySlice, course)
			}
		}
		return nonEmptySlice
	}
	return []string{}
}

func FilteredSliceIsNotEmpty(stringSlice []string) bool {
	newSlice := FilterEmptyStringsFromSlice(stringSlice)
	return (len(newSlice) > 0)
}

func StringArraySortUniqueLowercase(arrayToSort []string) []string {
	slices.Sort(arrayToSort)
	arrayToSort = slices.Compact(arrayToSort)
	lowerSlice := make([]string, 0, len(arrayToSort))
	for _, s := range arrayToSort {
		lowerSlice = append(lowerSlice, strings.ToLower(s))
	}
	return lowerSlice
}
