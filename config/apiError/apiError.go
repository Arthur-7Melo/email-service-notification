package apierror

import (
	"net/http"
)

type ApiError struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewApiError(message string, code int, cause []Cause) *ApiError {
	return &ApiError{
		Message: message,
		Code:    code,
		Causes:  cause,
	}
}

func (a *ApiError) Error() string {
	return a.Message
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, cause []Cause) *ApiError{
	return &ApiError{
		Message: message,
		Code: http.StatusBadRequest,
		Causes: cause,
	}
}

func NewInternalServerError(message string) *ApiError{
	return &ApiError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ApiError{
	return &ApiError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}