package formatter

type Response struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	TraceID    string      `json:"traceId,omitempty"`
	StatusCode *int        `json:"statusCodeClient,omitempty"`
	ErrorList  interface{} `json:"errorList,omitempty"`
}

func NewSuccessResponse(status Status, data interface{}) *Response {
	return &Response{
		Status:  status.String(),
		Message: "success",
		Data:    data,
	}
}

func NewErrorResponse(status Status, message string, id string) *Response {
	return &Response{
		Status:  status.String(),
		Message: message,
		TraceID: id,
	}
}

func NewErrorResponseList(status Status, message string, id string, err interface{}) *Response {
	return &Response{
		Status:    status.String(),
		Message:   message,
		TraceID:   id,
		ErrorList: err,
	}
}
