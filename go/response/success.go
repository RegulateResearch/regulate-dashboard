package response

type SuccessResponse[T any] struct {
	Message string
	Data    T
}

func NewSuccessResponse[T any](data T, message string) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    data,
		Message: message,
	}
}
