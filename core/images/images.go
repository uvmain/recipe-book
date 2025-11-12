package images

import (
	"bytes"
	"image"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"recipebook/core/config"
	"time"

	"github.com/gen2brain/webp"
	"github.com/nfnt/resize"
)

func UploadImage(file multipart.File, filename string) error {
	log.Printf("Uploading: %s", filename)

	filePath := filepath.Join(config.ImagesDirectory, filename)
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

func GetImageByFilename(filename string) ([]byte, time.Time, error) {
	filePath, _ := filepath.Abs(filepath.Join(config.ImagesDirectory, filename))

	info, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Image file does not exist: %s:  %s", filePath, err)
		return nil, time.Now(), err
	}

	modTime := info.ModTime()

	blob, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading image for filename %s: %s", filename, err)
		return nil, time.Now(), err
	}
	return blob, modTime, nil
}

func ResizeImageBytes(imgBytes []byte, size int, quality int) []byte {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Printf("Error decoding image: %s", err)
		return nil
	}

	if size > 0 {
		img = resize.Thumbnail(uint(size), uint(size), img, resize.Lanczos3)
	}

	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, webp.Options{Quality: quality}); err != nil {
		log.Printf("Error encoding image: %s", err)
		return nil
	}

	return buf.Bytes()
}

func DeleteImageByFilename(filename string) error {
	filePath, _ := filepath.Abs(filepath.Join(config.ImagesDirectory, filename))
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
