package models

// LoginResponse Devuelve un token para el login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
