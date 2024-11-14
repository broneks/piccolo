package backblaze

import (
	"context"
	"io"
	"piccolo/api/types"
)

func (b *BackblazeClient) UploadFile(ctx context.Context, fileUpload types.FileUpload) (string, error) {
	name := newObjectName(fileUpload.Filename, fileUpload.UserId)
	obj := b.bucket.Object(name)

	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, *fileUpload.File); err != nil {
		w.Close()
		return "", err
	}

	location := obj.URL()

	return location, w.Close()
}
