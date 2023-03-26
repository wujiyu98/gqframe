package database

import (
	"github.com/wujiyu98/gqframe/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := config.V.GetString("dsn")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
