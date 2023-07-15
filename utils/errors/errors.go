package errors

import "fmt"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.Code, e.Err)
}

func (e APIError) ErrorCode() int {
	return e.Code
}
