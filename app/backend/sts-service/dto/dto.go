package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ClientCredentialsRequest struct {
	ClientName string `json:"client_name,omitempty"`
}

func (c ClientCredentialsRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ClientName, validation.Required),
	)
}

type ClientCredentials struct {
	ClientName   string `json:"client_name,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}
