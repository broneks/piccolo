package upload

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func uploadFile(file *multipart.FileHeader) error {
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
