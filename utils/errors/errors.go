package errors

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return e.Message
}

func (e APIError) ErrorCode() int {
	return e.Code
}
