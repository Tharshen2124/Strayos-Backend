package main

import (
	"example/main/Controllers/StrayPetsController"
	"example/main/Controllers/UserController"
	"example/main/Middleware"
	"example/main/utils"
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
	frontendURL := utils.GetEnv("FRONTEND_URL")
	
	// routes
	r.Handle("/signup", middleware.Guest(http.HandlerFunc(UserController.LegacySignupUser))).Methods("POST")
	r.Handle("/login", middleware.Guest(http.HandlerFunc(UserController.LegacyLoginUser))).Methods("POST")
	r.Handle("/google-oauth/login", middleware.Guest(http.HandlerFunc(UserController.SignupWithGoogleOAuth))).Methods("POST")
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Index))).Methods("GET")
	r.Handle("/stray-pets", middleware.Auth(http.HandlerFunc(StrayPetsController.Create))).Methods("POST")

	// custom CORS settings
	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{frontendURL}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
        handlers.AllowCredentials(),
    )

	log.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", corsHandler(setCustomHeaders(r)))
}