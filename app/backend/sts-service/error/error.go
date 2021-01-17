package error

import (
	"net/http"
	"runtime/debug"
)

type IWithMessageAndStatusCode interface {
	Status() int
	Error() string
	Cause() string
	StackTrace() string
}

type ErrorSt struct {
	status     int
	msg        string
	cause      error
	stackTrace string
}

func (e ErrorSt) Status() int {
	return e.status
}

func (e ErrorSt) Error() string {
	return e.msg
}

func (e ErrorSt) Cause() string {
	if e.cause != nil {
		return e.cause.Error()
	} else {
		return ""
	}
}

func (e *ErrorSt) StackTrace() string {
	return e.stackTrace
}

type GenericError struct {
	ErrorSt
}

func (e GenericError) Code() int { return http.StatusInternalServerError }

func NewGenericError(msg string, cause error) error {
	return &GenericError{ErrorSt{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}

type NotFoundError struct {
	GenericError
}

func (e NotFoundError) Code() int { return http.StatusNotFound }

func NewNotFoundError(msg string) error {
	return &NotFoundError{GenericError{ErrorSt{msg: msg, stackTrace: string(debug.Stack())}}}
}

type AuthError struct {
	GenericError
}

func (e AuthError) Code() int { return http.StatusUnauthorized }

func NewAuthError(msg string, cause error) error {
	return &AuthError{GenericError{ErrorSt{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}}
}

type BadRequestError struct {
	GenericError
}

func (e BadRequestError) Code() int { return http.StatusBadRequest }

func NewBadRequestError(msg string, cause error) error {
	return &BadRequestError{GenericError{ErrorSt{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}}
}
