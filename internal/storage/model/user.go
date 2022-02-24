package model

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
	Active       bool   `json:"active"`
}
