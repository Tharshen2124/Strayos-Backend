package main

import (
	"example/main/Controllers/StrayPetsController"
	"example/main/Controllers/UserController"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	
	// routes
	r.HandleFunc("/signup", UserController.LegacySignupUser).Methods("POST")
	r.HandleFunc("/login", UserController.LegacyLoginUser).Methods("POST")
	r.HandleFunc("/auth/login", UserController.SignupWithGoogleOAuth).Methods("GET")
	r.HandleFunc("/auth/callback", UserController.CallBack).Methods("GET")
	r.HandleFunc("/stray-pets", StrayPetsController.Index).Methods("GET")
	r.HandleFunc("/stray-pets", StrayPetsController.Create).Methods("POST")

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8000", handler)
	fmt.Println("Server is running on localhost:8000")
}