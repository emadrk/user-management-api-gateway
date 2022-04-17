package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"usermanagement/pkg/models"
	"usermanagement/pkg/utils"

	"github.com/gorilla/mux"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Println("error whiel parsing")
	}
	userDetails, _ := models.GetUserById(Id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Println("Error wjhile parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.Address != "" {
		userDetails.Address = updateUser.Address
	}
	if updateUser.Domain != "" {
		userDetails.Domain = updateUser.Domain
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
