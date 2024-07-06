package log

import "fmt"

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewError(message string, code int) error {
	return &MyError{Message: message, Code: code}
}
