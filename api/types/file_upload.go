package types

import "mime/multipart"

type FileUpload struct {
	File     *multipart.File
	Filename string
	FileSize int32
	UserId   string
}
