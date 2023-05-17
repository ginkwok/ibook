package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/ginkwok/ibook/config"
	"github.com/ginkwok/ibook/model"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func init() {
	host := viper.GetString("db.mysql.host")
	port := viper.GetString("db.mysql.port")
	user := viper.GetString("db.mysql.username")
	pass := viper.GetString("db.mysql.password")
	database := viper.GetString("db.mysql.database")

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

	err = db.AutoMigrate(&model.Reservation{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
