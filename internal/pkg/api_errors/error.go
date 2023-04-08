package api_errors

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAPIError(message string, code int) *APIError {
	return &APIError{message, code}
}
