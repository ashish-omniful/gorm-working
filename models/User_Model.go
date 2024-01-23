package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
