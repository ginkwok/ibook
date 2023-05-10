package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user_table"
}
