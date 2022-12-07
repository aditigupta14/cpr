package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"Parking.com/DAL"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	users := DAL.GetUser()

	json.NewEncoder(w).Encode(users)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	_id := params["_id"]

	DAL.DeleteUserByID(_id)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	_id := params["_id"]

	var user DAL.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)

	fmt.Println(err)

	fmt.Println(user)

	DAL.UpdateUserByID(_id, user)

}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user DAL.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)

	fmt.Println(err)

	fmt.Println(user)

	DAL.AddUser(user)

	w.WriteHeader(http.StatusCreated) // 201
}

func main() {

	Router := mux.NewRouter()
	Router.HandleFunc("/user", GetUser).Methods("GET")
	Router.HandleFunc("/user", AddUser).Methods("POST")
	Router.HandleFunc("/user/{_id}", UpdateUser).Methods("PUT")
	Router.HandleFunc("/user/{_id}", DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", Router)

}
