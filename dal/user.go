package dal

import (
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func (d *dal) CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *dal) GetUserByName(db *gorm.DB, username string) (*model.User, error) {
	var user *model.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *dal) CheckUser(db *gorm.DB, username string, password string) (bool, error) {
	var user *model.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}
