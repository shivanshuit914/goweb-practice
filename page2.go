package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Email  string
	Active bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))
	users := []User{
		{"Test 1", "test1@test.com", true},
		{"Test 2", "test2@test.com", true},
		{"Test 3", "test3@test.com", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Users []User }{users})
	})

	http.ListenAndServe(":8080", nil)
}
