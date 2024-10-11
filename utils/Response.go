package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseType struct {
	Message string `json:"message"`
	Error any `json:"error"`
}

func OkResponse(data any, message string, w http.ResponseWriter) {
	response := ResponseType{
		Message: message,
		Error: data,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreatedResponse(data any, message string, w http.ResponseWriter) {
	response := ResponseType{
		Message: message,
		Error: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func BadResponse(errorData any, errorMessage string, w http.ResponseWriter) {
	errorResponse := ResponseType{
		Message: errorMessage,
		Error: errorData,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}