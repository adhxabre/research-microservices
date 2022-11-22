package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []Users{
	{
		Name:     "Admin",
		Email:    "admin@mail.com",
		Password: "admin123",
	},
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	var userData Users
	var isGetUser = false

	for i, user := range users {
		if i == index {
			isGetUser = true
			userData = user
		}
	}

	if isGetUser == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Users

	json.NewDecoder(r.Body).Decode(&data)

	users = append(users, data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	
	var data Users
	var isGetUser = false

	json.NewDecoder(r.Body).Decode(&data)

	for i, _ := range users {
		if i  == index {
			isGetUser = true
			users[i] = data
		}
	}

	if isGetUser == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	var isGetUser = false
	var idx = 0

	for i, _ := range users {
		if i == idx {
			isGetUser = true
			idx = index
		}
	}

	if isGetUser == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID not found")
		return
	}

	users = append(users[:index], users[index+1:]...)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ID successfully deleted")
}