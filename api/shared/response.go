package shared

type SuccessRes struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
