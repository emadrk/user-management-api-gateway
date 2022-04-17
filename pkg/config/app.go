package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectToDB() {
	d, err := gorm.Open("mysql", "root:Emmurules.1#@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
