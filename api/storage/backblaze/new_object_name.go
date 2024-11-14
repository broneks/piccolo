package backblaze

import "fmt"

func newObjectName(filename, userId string) string {
	return fmt.Sprintf("%s/%s", userId, filename)
}
