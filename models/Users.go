package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"default:'email'"`
	Password string `json:"password" gorm:"default:'1234'"`
}
