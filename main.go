package main

import (
	"example/main/Controllers/UserController"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/login", UserController.LoginUser).Methods("POST")
	r.HandleFunc("/signup", UserController.SignupUser).Methods("POST")
	r.HandleFunc("/user", UserController.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{userId}", UserController.GetUser).Methods("GET")

	http.ListenAndServe(":8000", r)
	fmt.Println("Server is running on localhost:8000")
}