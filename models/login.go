package models

//Login had the token which is returned with the login
type Login struct {
	Token string `json:"token,omitempty"`
}
