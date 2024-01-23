package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string `json:"name" gorm:"default:'user'"`
	Email    string `json:"email" gorm:"unique,size:255,not null"`
	Password string `json:"password" gorm:"not null"`
}

// not null -> field cannot remain empty
// unique -> each field should be distinct
// size:255 -> size of the field
// default:'some_value' -> determines default value of field
// column:'col_name' -> determines column name for that field
