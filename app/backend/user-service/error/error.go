package error

import "net/http"

type GenericError struct {
	msg string // description of error
}

func (e *GenericError) Error() string { return e.msg }
func (e *GenericError) Code() int     { return http.StatusInternalServerError }

func NewGenericError(msg string) error {
	return &GenericError{msg}
}

type NotFoundError struct {
	GenericError
}

func (e *NotFoundError) Error() string { return e.msg }
func (e *NotFoundError) Code() int     { return http.StatusNotFound }

func NewNotFoundError(msg string) error {
	return &NotFoundError{GenericError{msg: msg}}
}

type AuthError struct {
	GenericError
}

func (e *AuthError) Error() string { return e.msg }
func (e *AuthError) Code() int     { return http.StatusUnauthorized }

func NewAuthError(msg string) error {
	return &AuthError{GenericError{msg: msg}}
}

type BadRequestError struct {
	GenericError
}

func (e *BadRequestError) Error() string { return e.msg }
func (e *BadRequestError) Code() int     { return http.StatusBadRequest }

func NewBadRequestError(msg string) error {
	return &BadRequestError{GenericError{msg: msg}}
}
