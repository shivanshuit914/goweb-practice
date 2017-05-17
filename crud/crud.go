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
		tmpl := template.Must(template.ParseFiles("update.html"))
		user := User{}
		id := req.URL.Query().Get("id")
		var dbname, dbemail string
		var dbid int
		err := db.QueryRow("SELECT id, name, email FROM users where id = ?", id).Scan(&dbid, &dbname, &dbemail)
		if err != nil {
			http.Error(res, "Server error, unable to delete your account.", 500)
			return
		}

		user = User{Id: dbid, Name: dbname, Email: dbemail}
		tmpl.Execute(res, user)

		return
	}

	name := req.FormValue("name")
	email := req.FormValue("email")
	id := req.FormValue("id")

	_, err = db.Exec("UPDATE users set name = ?, email = ? WHERE id = ?", name, email, id)

	if err != nil {
		http.Error(res, "Server error, unable to create your account."+name, 500)
		return
	}

	res.Write([]byte("User updated!"))

}

func deletePage(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		http.Error(res, "Server error, unable to delete your account.", 500)
		return
	}
	res.Write([]byte("User deleted!"))
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
	http.HandleFunc("/update/", updatePage)
	http.HandleFunc("/delete/", deletePage)
	http.ListenAndServe(":8080", nil)
}
