package models

import (
	"log"
	"usermanagement/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Domain  string `json:"domain"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	log.Println("inside model createuser func")
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}
