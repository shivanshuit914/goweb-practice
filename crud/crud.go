package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type User struct {
	Id    int
	Name  string
	Email string
}

func createPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "create.html")
		return
	}

	name := req.FormValue("name")
	email := req.FormValue("email")

	_, err = db.Exec("INSERT INTO users(name, email) VALUES(?, ?)", name, email)

	if err != nil {
		http.Error(res, "Server error, unable to create your account."+name, 500)
		return
	}

	res.Write([]byte("User created!"))
}

func updatePage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		// id := req.FormValue("id")
		http.ServeFile(res, req, "create.html")
		return
	}
}

func deletePage(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		http.Error(res, "Server error, unable to delete your account.", 500)
		return
	}
	res.Write([]byte("User deleted!" + id))
}

func listPage(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("list.html"))
	rows, err := db.Query("SELECT id, name, email FROM users")
	users := []User{}
	for rows.Next() {
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		if err != nil {
			http.Error(res, "Server error, unable to list accounts.", 500)
			return
		}
		users = append(users, User{Id: id, Name: name, Email: email})
	}
	tmpl.Execute(res, struct{ Users []User }{users})
}

func main() {
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/create", createPage)
	http.HandleFunc("/list", listPage)
	http.HandleFunc("/update", updatePage)
	http.HandleFunc("/delete/", deletePage)
	http.ListenAndServe(":8080", nil)
}
