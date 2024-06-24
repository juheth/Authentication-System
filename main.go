package main

import (
	users "V1/internal/Users"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	GetUsers := users.GetUsers
	CreateUsers := users.CreateUsers

	r := mux.NewRouter()
	r.HandleFunc("/Users", GetUsers).Methods("GET")
	r.HandleFunc("/Users", CreateUsers).Methods("POST")
	r.HandleFunc("/Users/{id}", users.GetUser).Methods("GET")
	r.HandleFunc("/Users/{id}", users.Update).Methods("PUT")
	r.HandleFunc("/Users/{id}", users.DeleteTask).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
