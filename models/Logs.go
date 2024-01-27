package models

type Log struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}
