package UserController

import (
	"encoding/json"
	"example/main/DB"
	"example/main/Models"
	"example/main/utils"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ValidationError struct {
	Key   string
	Error string
}

func LegacySignupUser(w http.ResponseWriter, request *http.Request) {
	var user Models.User

	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
	if jsonDecoderError != nil {
		utils.BadResponse(jsonDecoderError, w)
		return
	}

	validate := utils.GetValidator()
	rules := map[string]string{
		"Username": "required",
		"Email":    "required",
		"Password": "required",
	}
	validate.RegisterStructValidationMapRules(rules, Models.User{})
	if validationErrors := validate.Struct(user); validationErrors != nil {
		utils.HandleValidationError(validationErrors, w)
		return
	}

	user.Password, _ = utils.HashPassword(user.Password)

	db.Create(&user)

	signedJwtToken, jwtTokenError := utils.CreateToken(user)
	if jwtTokenError != nil {
		log.Printf("Error during token creation: %v", jwtTokenError)
		utils.BadResponse(jsonDecoderError, w)
	}

	utils.AuthOkResponse(signedJwtToken, w)
}

func LegacyLoginUser(w http.ResponseWriter, request *http.Request) {
	var user Models.User
	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
	if jsonDecoderError != nil {
		utils.BadResponse(jsonDecoderError, w)
		return
	}

	validate := utils.GetValidator()
	rules := map[string]string{
		"Email":    "required",
		"Password": "required",
	}

	validate.RegisterStructValidationMapRules(rules, Models.User{})

	if validationErrors := validate.Struct(user); validationErrors != nil {
		utils.HandleValidationError(validationErrors, w)
		return
	}

	requestPassword := user.Password

	value := db.Where("email = ?", user.Email).Find(&user).Scan(&user)
	if value.Error != nil {
		utils.BadResponse(value.Error, w)
		return
	}

	comparePasswordHashError := utils.CheckPasswordHash(requestPassword, user.Password)

	if comparePasswordHashError != nil {
		utils.BadResponse("Password doesnt match", w)
		return
	}
	
	signedJwtToken, jwtTokenError := utils.CreateToken(user)
	if jwtTokenError != nil {
		utils.BadResponse(jsonDecoderError, w)
	}

	utils.AuthOkResponse(signedJwtToken, w)
}


func TestMethod(w http.ResponseWriter, request *http.Request) {
	signedJwtToken := "dog"

	utils.AuthOkResponse(signedJwtToken, w)
}