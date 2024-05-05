package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Posts []Post `json:"posts"`
}

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
