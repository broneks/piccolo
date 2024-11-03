package upload

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"piccolo/api/storage/wasabi"
)

func uploadFile(file *multipart.FileHeader) error {
	log.Printf("uploading file: %s", file.Filename)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uploader := wasabi.NewUploader()

	result, err := uploader.UploadFile(file.Filename, src)
	if err != nil {
		log.Println("Error uploading file:", err)
		return err
	}

	log.Printf("File uploaded successfully: %s\n", result.Location)

	return nil
}

func uploadFileLocal(file *multipart.FileHeader) error {
	log.Printf("uploading file: %s", file.Filename)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join("data", filepath.Base(file.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
