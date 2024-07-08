package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	users "github.com/juheth/Registration-Form.git/internal/Users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var db *sql.DB

func connectionDB() {
	// connection to the database
	connectionDB := "root:root@tcp(127.0.0.1:3306)/juheth"
	var err error
	db, err = sql.Open("mysql", connectionDB)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Established connection")
}

func loginFormHandler(w http.ResponseWriter, r *http.Request) {
	// Form HTML
	tpl :=
		`<!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <title>Login</title>
    </head>
    <body>
        <h2>Login</h2>
        <form action="/login" method="post">
            <label for="username">Username:</label><br>
            <input type="text" id="username" name="username"><br><br>
            <label for="lastname">Lastname:</label><br>
            <input type="text" id="lastname" name="lastname"><br><br>
            <input type="submit" value="Login">
        </form>
    </body>
    </html>`

	// como dice el dicho si funciona no lo toques (john me explica)
	t, err := template.New("login-form").Parse(tpl)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// solicitud POST para el login
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al analizar los datos del formulario", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	lastname := r.Form.Get("lastname")

	// Consulta a la base de datos para verificar el usuario
	Verify := "SELECT username FROM Users WHERE username = ? AND lastname = ? LIMIT 1"
	var result string
	err = db.QueryRow(Verify, username, lastname).Scan(&result)

	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "User not found")
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	default:
		fmt.Fprintf(w, "Welcome, %s!", result)
	}
}

func main() {
	connectionDB()

	r := mux.NewRouter()

	r.HandleFunc("/", loginFormHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Usuarios
	r.HandleFunc("/users", users.GetUsers).Methods("GET")
	r.HandleFunc("/users", users.CreateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", users.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", users.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", users.DeleteUser).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Servidor iniciado en :8080")
	log.Fatal(srv.ListenAndServe())
}
