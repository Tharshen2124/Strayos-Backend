package UserController

import (
	"encoding/json"
	"example/main/DB"
	"example/main/utils"
	"fmt"
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

func SignupUser(w http.ResponseWriter, request *http.Request) {
	var user DB.User
	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
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
		"Username": "required",
		"Email":  "required",
		"Password": "required",
	}

	validate.RegisterStructValidationMapRules(rules, DB.User{})

	if validationErrors := validate.Struct(user); validationErrors != nil {
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

	user.Password, _ = utils.HashPassword(user.Password)

	db.Create(&user)
	
	response := Response{
		Message: "Successfully registered user!",
		Data: user,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, request *http.Request) {
	var user DB.User
	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
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
		"Email":  "required",
		"Password": "required",
	}
	
	validate.RegisterStructValidationMapRules(rules, DB.User{})

	if validationErrors := validate.Struct(user); validationErrors != nil {
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

	requestPassword := user.Password

	value := db.Where("email = ?", user.Email).Find(&user).Scan(&user)
	if value.Error != nil {
		fmt.Println(value.Error)
		return
	}

	comparePasswordHashError := utils.CheckPasswordHash(requestPassword, user.Password)
	
	if comparePasswordHashError!= nil {
		errorResponse := ErrorResponse{
			Message: "An error occured",
			Error: "Password does not match",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := Response{
		Message: "Succesfully logged w user!",
		Data: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}