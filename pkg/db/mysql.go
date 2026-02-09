package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := viper.GetString("db.dsn")
	db, _ := gorm.Open(mysql.Open(dsn))
	DB = db
}
