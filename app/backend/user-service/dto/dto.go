package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Error struct {
	Message    string `json:"message,omitempty"`
	Cause      string `json:"cause,omitempty"`
	StackTrace string `json:"stacktrace,omitempty"`
}

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Firstname, validation.Required),
		validation.Field(&u.Lastname, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 20)),
	)
}

type UserUpdate struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Token     string `json:"auth_token,omitempty"`
}

type LoginData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (l LoginData) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 20)),
	)
}

type RecoverLoginData struct {
	Email string `json:"email,omitempty"`
}

type RecoverLoginDataResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
