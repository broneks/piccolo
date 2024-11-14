package upload

import (
	"context"
	"fmt"
	"mime/multipart"
	"piccolo/api/model"
	"piccolo/api/types"

	"github.com/jackc/pgx/v5/pgtype"
)

func (m *UploadModule) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader, userId string) error {
	var err error

	var filename = fileHeader.Filename
	var fileSize = int(fileHeader.Size)
	var contentType = fileHeader.Header.Get("Content-Type")

	m.server.Logger.Debug(fmt.Sprintf("uploading file: %s", filename))

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	location, err := m.server.ObjectStorage.UploadFile(ctx, types.FileUpload{
		File:     &file,
		Filename: filename,
		UserId:   userId,
	})
	if err != nil {
		return err
	}

	m.server.Logger.Debug(fmt.Sprintf("File uploaded successfully: %s\n", location))

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
