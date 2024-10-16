package StrayPetsController

import (
	"encoding/json"
	"example/main/DB"
	"example/main/Models"
	"example/main/utils"
	"fmt"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error any `json:"error"`
}

type ValidationError struct {
	Key string
	Error string
}

func Index(w http.ResponseWriter, request *http.Request) {
    var strayPets []Models.StrayPet
    db := DB.DBConnect()

    result := db.Find(&strayPets)
    if result.Error != nil {
        log.Printf("Error fetching stray pets: %v", result.Error)
        return
    }

    response := Response{
		Data: strayPets,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func Create(w http.ResponseWriter, request *http.Request) {
    var strayPet Models.StrayPet
    db := DB.DBConnect()

    jsonDecoderError := json.NewDecoder(request.Body).Decode(&strayPet)
    if jsonDecoderError != nil {
        errorResponse := ErrorResponse{
            Message: "An error occured during decoding",
            Error: jsonDecoderError,
        }
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(errorResponse)
        return
    }

    validate := utils.GetValidator()
    rules := map[string]string{
        "Animal": "required",
        "UserId": "required",
        "Status": "required",
        "Latitude": "required",
        "Longitude": "required",
    }

    validate.RegisterStructValidationMapRules(rules, Models.StrayPet{})
    if validationErrors := validate.Struct(strayPet); validationErrors != nil {
		errorMap := make(map[string]interface{})
        for _, validationError := range validationErrors.(validator.ValidationErrors) {
			validationErrorValue := fmt.Sprintf("This Field with validation '%s' has failed",validationError.ActualTag())
			errorMap[validationError.Field()] = validationErrorValue 
        }
		errorResponse := ErrorResponse{
			Message: "An error occured during validation",
			Error: errorMap,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

    db.Create(&strayPet)

     // Preload the User data
     var savedStrayPet Models.StrayPet
     db.Preload("User").First(&savedStrayPet, strayPet.StrayPetId)
 
     response := Response{
         Data: savedStrayPet,
     }
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusCreated)
     json.NewEncoder(w).Encode(response)
}
