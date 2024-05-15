package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    string `json:"id" gorm:"primarykey"`
	Name  string `json:"name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Posts []Post `json:"posts"`
}

func NewUser(id, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
