package response

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewSuccessResponse[T any](data T, message string) SuccessResponse[T] {
	return SuccessResponse[T]{
		Data:    data,
		Message: message,
	}
}
