package models

type AuthData struct {
	ID       string `json:"id,omitempty"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
}
