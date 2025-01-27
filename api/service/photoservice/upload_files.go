package photoservice

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"piccolo/api/model"
	"piccolo/api/types"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
)

func (svc *PhotoService) handleFileUpload(ctx context.Context, fileHeader *multipart.FileHeader, userId string, resultCh chan<- model.Photo) {
	fileSize := int32(fileHeader.Size)

	file, err := fileHeader.Open()
	if err != nil {
		slog.Error("error opening file", "err", err)
		return
	}
	defer file.Close()

	slog.Debug("uploading file", "filename", fileHeader.Filename)

	location, err := svc.server.ObjectStorage.UploadFile(ctx, types.FileUpload{
		File:     &file,
		Filename: fileHeader.Filename,
		FileSize: fileSize,
		UserId:   userId,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("error uploading file \"%s\"", fileHeader.Filename), "err", err)
		return
	}

	slog.Debug("file uploaded successfully", "location", location)

	photo := model.Photo{
		Location:    pgtype.Text{String: location, Valid: true},
		Filename:    pgtype.Text{String: fileHeader.Filename, Valid: true},
		FileSize:    pgtype.Int4{Int32: fileSize, Valid: true},
		ContentType: pgtype.Text{String: fileHeader.Header.Get("Content-Type"), Valid: true},
	}

	resultCh <- photo
}

// TODO: batch upload
func (svc *PhotoService) UploadFiles(ctx context.Context, fileHeaders []*multipart.FileHeader, userId string) ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var photos []model.Photo

	isUserFileStorageExceeded, err := svc.isUserFileStorageExceeded(ctx, userId, fileHeaders)
	if isUserFileStorageExceeded || err != nil {
		return []string{}, err
	}

	resultCh := make(chan model.Photo, len(fileHeaders))

	for _, fileHeader := range fileHeaders {
		if fileHeader == nil {
			slog.Debug("recieved nil fileHeader")
			continue
		}

		wg.Add(1)

		go func(fh *multipart.FileHeader) {
			defer wg.Done()
			svc.handleFileUpload(context.WithoutCancel(ctx), fileHeader, userId, resultCh)
		}(fileHeader)
	}

	wg.Wait()

	close(resultCh)

	for photo := range resultCh {
		mu.Lock()
		photos = append(photos, photo)
		mu.Unlock()
	}

	if len(photos) == 0 {
		return []string{}, nil
	}

	ids, err := svc.photoRepo.InsertMany(ctx, photos, userId)
	if err != nil {
		return []string{}, fmt.Errorf("failed to upload photos: %v", err)
	}

	return ids, nil
}
