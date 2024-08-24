package dto

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateError(status int, message string) Response[string] {
	return Response[string]{
		Status:  status,
		Message: message,
		Data:    "",
	}
}

func CreateSuccess[T any](status int, message string, data T) Response[T] {
	return Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CreateResponseErrorData(status int, message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
