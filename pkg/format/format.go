package format

type (
	BaseResponse struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
		Result  any    `json:"result,omitempty"`
		Errors  any    `json:"errors,omitempty"`
	}
)

func Success(data any) BaseResponse {
	return BaseResponse{
		Status: "SUCCESS",
		Result: data,
	}
}

// 400 - Bad Request
func BadRequest(errors any) BaseResponse {
	return BaseResponse{
		Status: "FAILED",
		Errors: errors,
	}
}

func Failed(message string) BaseResponse {
	return BaseResponse{
		Message: message,
	}
}
