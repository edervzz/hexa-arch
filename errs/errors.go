package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewErrorNotFound() *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "Item not found",
	}
}

func NewUnexpectedError() *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "Item not found",
	}
}

func NewErrorMinimalBalance() *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "At least 5000.00 for opening account.",
	}
}

func NewBadRequest(text string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: text,
	}
}
