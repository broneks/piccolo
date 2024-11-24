package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"piccolo/api/model"
	"piccolo/api/repo/photorepo"
	"piccolo/api/types"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
)

type PhotoService struct {
	server    *types.Server
	photoRepo *photorepo.PhotoRepo
}

func NewPhotoService(server *types.Server, photoRepo *photorepo.PhotoRepo) *PhotoService {
	return &PhotoService{
		server:    server,
		photoRepo: photoRepo,
	}
}

func (svc *PhotoService) handleFileUpload(ctx context.Context, fileHeader *multipart.FileHeader, userId string, resultCh chan<- model.Photo) {
	fileSize := int32(fileHeader.Size)

	file, err := fileHeader.Open()
	if err != nil {
		svc.server.Logger.Error(fmt.Sprintf("Error opening file %s: %v", fileHeader.Filename, err))
		return
	}
	defer file.Close()

	svc.server.Logger.Debug(fmt.Sprintf("Uploading file: %s", fileHeader.Filename))

	location, err := svc.server.ObjectStorage.UploadFile(ctx, types.FileUpload{
		File:     &file,
		Filename: fileHeader.Filename,
		FileSize: fileSize,
		UserId:   userId,
	})
	if err != nil {
		svc.server.Logger.Error(fmt.Sprintf("Error uploading file \"%s\": %v", fileHeader.Filename, err))
		return
	}

	svc.server.Logger.Debug(fmt.Sprintf("File uploaded successfully: %s\n", location))

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

	resultCh := make(chan model.Photo, len(fileHeaders))

	for _, fileHeader := range fileHeaders {
		if fileHeader == nil {
			svc.server.Logger.Debug("received nil fileHeader")
			continue
		}

		wg.Add(1)

		go func(fh *multipart.FileHeader) {
			defer wg.Done()
			svc.handleFileUpload(ctx, fileHeader, userId, resultCh)
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
		return nil, nil
	}

	ids, err := svc.photoRepo.InsertMany(ctx, photos, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to insert photos: %v", err)
	}

	return ids, nil
}
