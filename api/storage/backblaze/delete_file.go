package backblaze

import (
	"context"
	"fmt"
	"log/slog"
)

func (bc *BackblazeClient) DeleteFile(ctx context.Context, fileName, userId string) error {
	name := newObjectName(fileName, userId)
	obj := bc.bucket.Object(name)

	if obj == nil {
		return fmt.Errorf("File does not exist")
	}

	slog.Debug("deleting file", "bucket", bc.bucket.Name(), "name", name)

	return obj.Delete(ctx)
}
