package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage return message without Code
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewBadRequestError return error for bad request
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

// NewNotFoundError return error for not found
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError return error for unexpected
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// NewValidationError return error for validation
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

// NewAuthenticationError return error for authentication
func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

// NewAuthorizationError return error for authorization
func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
