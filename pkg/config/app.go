package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func connect() {
	var err error
	db, err = gorm.Open("mysql", "rhaqim:saveyourpity@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
