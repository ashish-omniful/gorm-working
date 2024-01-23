package models

type User struct {
	ID       string `json:"id"`
	Value    string `json:"value"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
