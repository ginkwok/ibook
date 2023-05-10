package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("fail to read config file:" + err.Error())
	}

	viper.BindEnv("db.mysql.host", "MYSQL_HOST")
	viper.BindEnv("db.mysql.port", "MYSQL_PORT")
	viper.BindEnv("db.mysql.username", "MYSQL_USERNAME")
	viper.BindEnv("db.mysql.password", "MYSQL_PASSWORD")
	viper.BindEnv("db.mysql.database", "MYSQL_DB")

	viper.BindEnv("jwt.key", "JWT_KEY")
}
