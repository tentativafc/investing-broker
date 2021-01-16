package error

import (
	"net/http"
	"runtime/debug"
)

type withMessageAndCause struct {
	msg        string
	cause      error
	stackTrace string
}

func (w *withMessageAndCause) Error() string {
	return w.msg
}

func (w *withMessageAndCause) Cause() string {
	if w.cause != nil {
		return w.cause.Error()
	} else {
		return ""
	}
}

func (w *withMessageAndCause) StackTrace() string {
	return w.stackTrace
}

type GenericError struct {
	withMessageAndCause
}

func (e *GenericError) Code() int { return http.StatusInternalServerError }

func NewGenericError(msg string, cause error) error {
	return &GenericError{withMessageAndCause{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}

type NotFoundError struct {
	GenericError
}

func (e *NotFoundError) Code() int { return http.StatusNotFound }

func NewNotFoundError(msg string) error {
	return &NotFoundError{GenericError{withMessageAndCause{msg: msg, stackTrace: string(debug.Stack())}}}
}

type AuthError struct {
	GenericError
}

func (e *AuthError) Code() int { return http.StatusUnauthorized }

func NewAuthError(msg string, cause error) error {
	return &AuthError{GenericError{withMessageAndCause{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}}
}

type BadRequestError struct {
	GenericError
}

func (e *BadRequestError) Code() int { return http.StatusBadRequest }

func NewBadRequestError(msg string, cause error) error {
	return &BadRequestError{GenericError{withMessageAndCause{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}}
}
