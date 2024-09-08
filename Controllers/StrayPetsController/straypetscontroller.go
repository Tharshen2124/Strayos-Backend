package StrayPetsController

import (
	"example/main/DB"
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

func Index(w http.ResponseWriter, request *http.Request) {
	var pets []DB.StrayPet
	db := DB.DBConnect()
	db.Find(&pets)

	response := Response{
		Data: pets,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Show(w http.ResponseWriter, request *http.Request) {
	
}

func Create(w http.ResponseWriter, request *http.Request) {

}

func Update(w http.ResponseWriter, request *http.Request) {

}

func Destroy(w http.ResponseWriter, request *http.Request) {

}
