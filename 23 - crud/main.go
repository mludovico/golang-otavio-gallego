package main

import (
	"crud/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/user", handlers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/user/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
