package photos

import "time"

type PhotoRes struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Filename    string    `json:"filename"`
	FileSize    int       `json:"fileSize"`
	Url         string    `json:"url"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
}
