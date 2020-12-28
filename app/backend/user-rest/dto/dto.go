package dto

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
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
