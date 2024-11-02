package wasabi

import (
	"context"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type WasabiUploader struct {
	Config   *WasabiConfig
	uploader *manager.Uploader
}

func (w *WasabiUploader) UploadFile(filename string, file multipart.File) (*manager.UploadOutput, error) {
	return w.uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(w.Config.Bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
}

func NewUploader() *WasabiUploader {
	client, err := newClient()
	if err != nil {
		return nil
	}

	wasabiConfig := &WasabiConfig{
		Bucket: os.Getenv("WASABI_BUCKET_NAME"),
		Region: os.Getenv("WASABI_BUCKET_REGION"),
	}

	return &WasabiUploader{
		Config:   wasabiConfig,
		uploader: manager.NewUploader(client),
	}
}
