package utils

import (
	"encoding/json"
	"net/http"
)

type BadResponseType struct {
	Message string `json:"message"`
	Error any `json:"error"`
}

type ResponseType struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type AuthResponseType struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

func OkResponse(data any, message string, w http.ResponseWriter) {
	response := ResponseType{
		Message: message,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AuthOkResponse(token string, w http.ResponseWriter) {
	response := AuthResponseType{
		Message: "Successfully signed in!",
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


func CreatedResponse(data any, message string, w http.ResponseWriter) {
	response := ResponseType{
		Message: message,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func BadResponse(errorData any, w http.ResponseWriter) {
	errorResponse := BadResponseType{
		Message: "Error occured",
		Error: errorData,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}