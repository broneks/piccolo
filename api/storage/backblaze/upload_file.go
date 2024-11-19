package backblaze

import (
	"context"
	"io"
	"piccolo/api/types"
)

func getConcurrentUploads(bytes int32) int {
	if bytes == 0 {
		return 1
	}

	gb := float64(bytes) / (1024 * 1024 * 1024)

	if gb >= 0.75 {
		return 4 // heuristic
	}

	return 1
}

func (bc *BackblazeClient) UploadFile(ctx context.Context, fileUpload types.FileUpload) (string, error) {
	name := newObjectName(fileUpload.Filename, fileUpload.UserId)
	obj := bc.bucket.Object(name)

	w := obj.NewWriter(ctx)

	w.ConcurrentUploads = getConcurrentUploads(fileUpload.FileSize)

	if _, err := io.Copy(w, *fileUpload.File); err != nil {
		w.Close()
		return "", err
	}

	location := obj.URL()

	return location, w.Close()
}
