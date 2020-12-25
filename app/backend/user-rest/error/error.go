package error

type GenericError struct {
	msg string // description of error
}

func (e *GenericError) Error() string { return e.msg }

func NewGenericError(msg string) error {
	return &GenericError{msg}
}

type NotFoundError struct {
	GenericError
}

func (e *NotFoundError) Error() string { return e.msg }

func NewNotFountError(msg string) error {
	return &NotFoundError{GenericError{msg: msg}}
}

type AuthError struct {
	GenericError
}

func (e *AuthError) Error() string { return e.msg }

func NewAuthError(msg string) error {
	return &AuthError{GenericError{msg: msg}}
}
