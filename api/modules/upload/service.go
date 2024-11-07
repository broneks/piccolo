package upload

import (
	"context"
	"log"
	"mime/multipart"
	"piccolo/api/model"
)

func (m *UploadModule) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
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

	location, err := m.server.ObjectStorage.UploadFile(ctx, filename, src)
	if err != nil {
		log.Println("Error uploading file:", err)
		return err
	}

	log.Printf("File uploaded successfully: %s\n", location)

	photo := model.Photo{
		Location:    location,
		Filename:    filename,
		FileSize:    fileSize,
		ContentType: contentType,
	}

	err = m.photoRepo.InsertOne(ctx, photo)
	if err != nil {
		return err
	}

	return nil
}
