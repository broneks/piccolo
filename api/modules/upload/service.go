package upload

import (
	"context"
	"log"
	"mime/multipart"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5/pgtype"
)

func (m *UploadModule) UploadFile(ctx context.Context, file *multipart.FileHeader, userId string) error {
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
		return err
	}

	log.Printf("File uploaded successfully: %s\n", location)

	photo := model.Photo{
		UserId:      pgtype.Text{String: userId, Valid: true},
		Location:    pgtype.Text{String: location, Valid: true},
		Filename:    pgtype.Text{String: filename, Valid: true},
		FileSize:    pgtype.Int4{Int32: int32(fileSize), Valid: true},
		ContentType: pgtype.Text{String: contentType, Valid: true},
	}

	err = m.photoRepo.InsertOne(ctx, photo)
	if err != nil {
		return err
	}

	return nil
}
