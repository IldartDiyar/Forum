package project_error

import "net/http"

func NewUserError(msg string, status int) error {
	return &UserError{msg, status}
}

type UserError struct {
	msg    string
	status int
}

func (e *UserError) Error() string {
	return e.msg
}

func (e *UserError) Status() int {
	return e.status
}

func NewInternalError(msg string) error {
	return &InternalError{msg}
}

type InternalError struct {
	msg string
}

func (e *InternalError) Error() string {
	return e.msg
}

type ServerError struct {
	Message string
	Status  int
}

func (e ServerError) Error() string {
	return e.Message
}

func NewServerError(message string) error {
	return ServerError{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}
