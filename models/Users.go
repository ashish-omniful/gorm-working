package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" gorm:"default:'1234'"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	fmt.Println("successfully created data")
	return nil
}

func (u *User) AfterUpdate(db *gorm.DB) (err error) {
	fmt.Println("successfully updated data")
	return nil
}
