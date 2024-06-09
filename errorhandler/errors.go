package errorhandler

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewErrorNotFound(msg string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewErrorBadRequest(msg string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NewErrorInternalServerError(msg string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func NewErrorUnauthorized(msg string) error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func NewErrorForbidden(msg string) error {
	return AppError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}
