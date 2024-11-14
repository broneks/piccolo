package wasabi

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// TODO: organize by folder that namespaces to user id
func (wc *WasabiClient) UploadFile(ctx context.Context, filename string, file multipart.File) (string, error) {
	result, err := wc.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(wc.config.bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
