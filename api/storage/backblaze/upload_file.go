package backblaze

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
)

func (b *BackblazeClient) UploadFile(ctx context.Context, file multipart.File, filename, userId string) (string, error) {
	dst := fmt.Sprintf("%s/%s", userId, filename)
	obj := b.bucket.Object(dst)

	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, file); err != nil {
		w.Close()
		return "", err
	}

	// TODO: get location
	return "", w.Close()
}
