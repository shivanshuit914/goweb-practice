package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Email     string   `json:"email,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []User

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person User
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func UpdateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var address *Address
	for _, item := range people {
		if item.ID == params["id"] {
			item.Firstname = params["firstname"]
			item.Email = params["email"]
			address.City = params["city"]
			address.State = params["state"]
			item.Address = address
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, User{ID: "1", Firstname: "User1", Email: "test@restapi.com", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, User{ID: "2", Firstname: "User2", Email: "test@restapi.com"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/people/", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", UpdateUserEndpoint).Methods("PUT")
	router.HandleFunc("/people/{id}", DeleteUserEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
