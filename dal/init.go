package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func GetDB(host string, port string, user string, pass string, database string) *gorm.DB {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect mysql db:" + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Room{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Seat{})
	if err != nil {
		panic(err)
	}

	return db
}
