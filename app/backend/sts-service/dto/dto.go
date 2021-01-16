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

type TokenRequest struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func (c TokenRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ClientId, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
	)
}

type TokenResponse struct {
	Token string `json:"token,omitempty"`
}

type ValidateTokenRequest struct {
	Token string `json:"token,omitempty"`
}

func (c ValidateTokenRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Token, validation.Required),
	)
}

type ValidateTokenResponse struct {
	Token      string `json:"token,omitempty"`
	ClientId   string `json:"client_id,omitempty"`
	ClientName string `json:"client_name,omitempty"`
}
