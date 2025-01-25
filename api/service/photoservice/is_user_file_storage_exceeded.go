package photoservice

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
)

func (svc *PhotoService) isUserFileStorageExceeded(ctx context.Context, userId string, fileHeadersToUpload []*multipart.FileHeader) (bool, error) {
	userFileStorage, err := svc.GetUserFileStorage(ctx, userId)
	if err != nil {
		slog.Error("failed to get user file storage when uploading photos", "err", err)

		return false, fmt.Errorf("failed to upload photos: %v", err)
	}

	if userFileStorage.UsedPercentage >= 100 {
		return true, fmt.Errorf("file storage limit exceeded: the user has already exceeded their file storage limit.")
	}

	var totalFileSizeMB float32

	for _, fileHeader := range fileHeadersToUpload {
		if fileHeader == nil {
			continue
		}

		totalFileSizeMB += float32(fileHeader.Size) / (1024 * 1024)
	}

	if (userFileStorage.UsedMB + totalFileSizeMB) > userFileStorage.TotalMB {
		return true, fmt.Errorf("file storage limit exceeded: uploading the photos will exceed the user's file storage limit.")
	}

	return false, nil
}
