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
		http.ServeFile(res, req, "list.html")
		return
	}
}

func listPage(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("list.html"))
	rows, err := db.Query("SELECT name, email FROM users")
	i := 0
	users := make([]User, 3)
	for rows.Next() {
		var name string
		var email string
		err = rows.Scan(&name, &email)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}
		users[i] = User{Name: name, Email: email}
		i++
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
	http.ListenAndServe(":8080", nil)
}
