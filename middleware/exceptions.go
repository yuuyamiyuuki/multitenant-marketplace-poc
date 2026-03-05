package middleware

import "fmt"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewBadRequest(message string, err error) *AppError {
	return &AppError{Code: 400, Message: message, Err: err}
}

func NewNotFound(message string, err error) *AppError {
	return &AppError{Code: 404, Message: message, Err: err}
}

func NewInternal(message string, err error) *AppError {
	return &AppError{Code: 500, Message: message, Err: err}
}

func NewConflict(message string, err error) *AppError {
	return &AppError{Code: 409, Message: message, Err: err}
}
