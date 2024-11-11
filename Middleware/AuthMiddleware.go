package middleware

import (
	"encoding/json"
	"example/main/Log"
	"example/main/utils"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		Log.Route(request)
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]

		_, err := utils.ParseJWTToken(tokenString)
 
		if err != nil {
			errorResponse := utils.BadResponseType{
				Message: "Error occured",
				Error: err,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		next.ServeHTTP(w, request)
	})
}

func Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		Log.Route(request)
		next.ServeHTTP(w, request)
	})
}
