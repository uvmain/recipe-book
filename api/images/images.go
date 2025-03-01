package images

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"recipebook/logic"
)

func UploadImage(file multipart.File, filename string) error {
	log.Printf("Uploading: %s", filename)

	filePath := filepath.Join(logic.ImagesDirectory, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Printf("failed to parse uploaded file: %v", err)
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		log.Printf("failed to write uploaded file: %v", err)
		return err
	}

	log.Printf("Uploaded file to %s", filePath)
	return nil
}

func GetImageByFilename(filename string) ([]byte, error) {
	filePath, _ := filepath.Abs(filepath.Join(logic.ImagesDirectory, filename))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("Image file does not exist: %s:  %s", filePath, err)
		return nil, err
	}
	blob, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading image for filename %s: %s", filename, err)
		return nil, err
	}
	return blob, nil
}

func DeleteImageByFilename(filename string) error {
	filePath, _ := filepath.Abs(filepath.Join(logic.ImagesDirectory, filename))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("Image file does not exist: %s:  %s", filePath, err)
		return err
	}
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
