package upload

import (
	"context"
	"log"
	"mime/multipart"
	"piccolo/api/model"
	"piccolo/api/storage/pg"
	"piccolo/api/storage/wasabi"
)

func uploadFile(file *multipart.FileHeader) error {
	var err error

	var filename = file.Filename
	var fileSize = int(file.Size)
	var contentType = file.Header.Get("Content-Type")

	log.Printf("uploading file: %s", filename)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uploader := wasabi.NewUploader(context.Background())

	result, err := uploader.UploadFile(context.Background(), filename, src)
	if err != nil {
		log.Println("Error uploading file:", err)
		return err
	}

	log.Printf("File uploaded successfully: %s\n", result.Location)

	db := pg.Client(context.Background())

	photo := model.Photo{
		Location:    result.Location,
		Filename:    filename,
		FileSize:    fileSize,
		ContentType: contentType,
	}

	err = db.InsertPhoto(context.Background(), photo)
	if err != nil {
		return err
	}

	return nil
}
