package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;size:255;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user_table"
}
