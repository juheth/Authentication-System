package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	login "github.com/juheth/Registration-Form.git/internal/Login"
	register "github.com/juheth/Registration-Form.git/internal/Register"
	users "github.com/juheth/Registration-Form.git/internal/Users"
)

var db *sql.DB

// connection to the database
func connectionDB() {
	connectionDB := "root:root@tcp(127.0.0.1:3306)/juheth"
	var err error
	db, err = sql.Open("mysql", connectionDB)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Established connection")
	login.SetDB(db)
	register.SetDB(db)
}

func main() {
	connectionDB()

	r := mux.NewRouter()

	// login
	r.HandleFunc("/", login.LoginFormHandler).Methods("GET")
	r.HandleFunc("/login", register.UserDoesNotExist).Methods("POST")

	// Register Users
	r.HandleFunc("/register", register.RegisterFormHandler).Methods("GET")
	r.HandleFunc("/register", register.RegisterHandler).Methods("POST")

	// Users
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
