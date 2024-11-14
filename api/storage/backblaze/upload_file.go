package backblaze

import (
	"context"
	"io"
	"mime/multipart"
)

func (b *BackblazeClient) UploadFile(ctx context.Context, file multipart.File, filename, userId string) (string, error) {
	name := newObjectName(filename, userId)
	obj := b.bucket.Object(name)

	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, file); err != nil {
		w.Close()
		return "", err
	}

	location := obj.URL()

	return location, w.Close()
}
