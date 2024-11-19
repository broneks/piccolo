package backblaze

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Backblaze/blazer/b2"
)

func (bc *BackblazeClient) doesObjectExist(ctx context.Context, filename, userId string) bool {
	if filename == "" {
		return false
	}

	iterator := bc.bucket.List(ctx, b2.ListPrefix(userId))

	for iterator.Next() {
		obj := iterator.Object()
		name := obj.Name()

		if strings.HasSuffix(name, filename) {
			return true
		}
	}

	if err := iterator.Err(); err != nil {
		slog.Debug(fmt.Sprintf("encountered a backblaze object iterator error: %v", err))
	}

	return false
}
