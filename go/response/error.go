package response

import (
	"frascati/exception"
)

type ErrorResponse[T any] struct {
	Message string
	Error   T
}

func NewExceptionResponse(message string, exc exception.Exception) ErrorResponse[string] {
	return ErrorResponse[string]{
		Message: message,
		Error:   exc.Error(),
	}
}

func NewErrorResponse[T any](message string, errObj T) ErrorResponse[T] {
	return ErrorResponse[T]{
		Message: message,
		Error:   errObj,
	}
}
