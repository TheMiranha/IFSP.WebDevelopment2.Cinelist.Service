package dtos

type RequestError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewRequestError(message string) *RequestError {
	return &RequestError{
		Success: false,
		Message: message,
	}
}
