package common

type DefaultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequestResponse() *DefaultResponse {
	return &DefaultResponse{
		"failed",
		"Bad request",
		"",
	}
}

func InternalServerErrorResponse(err string) *DefaultResponse {
	return &DefaultResponse{
		"failed",
		"Bad request",
		err,
	}
}

func NewSuccessResponse(data interface{}) *DefaultResponse {
	return &DefaultResponse{
		"success",
		"OK",
		data,
	}
}
