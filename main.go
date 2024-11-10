package main

import (
	"example/main/Controllers/StrayPetsController"
	"example/main/Controllers/UserController"
	"example/main/Middleware"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Middleware to add custom headers
func setCustomHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin-allow-popups")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	
	// routes
	r.Handle("/signup", middleware.Guest(http.HandlerFunc(UserController.LegacySignupUser))).Methods("POST")
	r.Handle("/login", middleware.Guest(http.HandlerFunc(UserController.LegacyLoginUser))).Methods("POST")
	r.Handle("/google-oauth/login", middleware.Guest(http.HandlerFunc(UserController.SignupWithGoogleOAuth))).Methods("POST")
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Index))).Methods("GET") // this works
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Create))).Methods("POST")
	r.Handle("/test-guest", middleware.Guest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Guest route worked!"))
	}))).Methods("POST")
	
	// custom CORS settings
	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", "https://strayos.vercel.app"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
        handlers.AllowCredentials(),
    )

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler(setCustomHeaders(r)))
}