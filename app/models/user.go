package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;size:255" json:"email"`
	Password string `gorm:"not null" json:"-"`
}
