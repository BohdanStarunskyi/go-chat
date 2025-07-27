package dto

type MapResponse map[string]any

type DefaultResponseWrapper[T any] struct {
	Data T `json:"data"`
}

func NewErrorResponse(err string) MapResponse {
	return MapResponse{"error": err}
}

func NewMessageResponse(message string) MapResponse {
	return MapResponse{"message": message}
}

func NewDataResponse(data any) MapResponse {
	return MapResponse{"data": data}
}
