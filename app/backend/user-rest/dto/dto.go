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
}

type LoginData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RecoverLoginData struct {
	Email string `json:"email,omitempty"`
}

type LoginResponse struct {
	Token    string       `json:"auth_token,omitempty"`
	UserData UserResponse `json:"user,omitempty"`
}

type RecoverLoginDataResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type ErrorResponse struct {
	Msg string `json:"message,omitempty"`
}