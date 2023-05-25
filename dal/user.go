package dal

import (
	"github.com/ginkwok/ibook/model"
)

func (d *dal) CreateUser(user *model.User) (*model.User, error) {
	err := d.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *dal) GetUserByName(username string) (*model.User, error) {
	var user *model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *dal) CheckUser(username string, password string) (bool, error) {
	var user *model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}
