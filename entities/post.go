package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	User    User   `gorm:"foreignkey:UserId; constraint:OnDelete:CASCADE"`
	UserId  uint   `gorm:"not null"`
}
