package backblaze

import "fmt"

// TODO group by year and month as well
func newObjectName(filename, userId string) string {
	return fmt.Sprintf("%s/%s", userId, filename)
}
