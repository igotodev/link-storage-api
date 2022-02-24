package model

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Active       bool   `json:"active"`
}
