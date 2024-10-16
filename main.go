package main

import (
	"example/main/Controllers/StrayPetsController"
	"example/main/Controllers/UserController"
	"example/main/Middleware"
	"fmt"
	"net/http"
    "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	// routes
	r.HandleFunc("/signup", UserController.LegacySignupUser).Methods("POST")
	r.HandleFunc("/login", UserController.LegacyLoginUser).Methods("POST")
	// r.HandleFunc("/auth/login", UserController.SignupWithGoogleOAuth).Methods("GET")
	// r.HandleFunc("/auth/callback", UserController.CallBack).Methods("GET")
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Index))).Methods("GET")
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Create))).Methods("POST")

	 // Define custom CORS settings
	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
        handlers.AllowCredentials(),
    )

	http.ListenAndServe(":8000", corsHandler(r))
	fmt.Println("Server is running on localhost:8000")
}