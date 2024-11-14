package types

import "mime/multipart"

type FileUpload struct {
	File     *multipart.File
	Filename string
	UserId   string
}
