package UserController

import (
	"context"
	"encoding/json"
	"example/main/DB"
	"example/main/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type ValidationError struct {
	Key string
	Error string
}

var oAuthConfig *oauth2.Config

func LegacySignupUser(w http.ResponseWriter, request *http.Request) {
	var user DB.User

	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
	if jsonDecoderError != nil {
		message := "An error occured during decoding"
		utils.BadResponse(jsonDecoderError, message, w)
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
		utils.HandleValidationError(validationErrors, w)
		return
	}

	user.Password, _ = utils.HashPassword(user.Password)

	db.Create(&user)
	
	message := "User successfully registered!"
	utils.OkResponse(user, message, w)
}

func SignupWithGoogleOAuth(w http.ResponseWriter, request *http.Request) {

	if err := godotenv.Load(); err != nil {
    	log.Fatal("Error loading .env file")
  	}

	oAuthConfig = &oauth2.Config{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8000/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	log.Printf("oAuthConfig: %v", oAuthConfig)

	url := oAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Printf("url :%v", url)

	http.Redirect(w, request, url, http.StatusTemporaryRedirect)
}

func CallBack(w http.ResponseWriter, request *http.Request) {
    code := request.URL.Query().Get("code")

	log.Printf("code: %v", code)

    if code == "" {
        http.Error(w, "No code in the request", http.StatusBadRequest)
        return
    }

    ctx := context.Background()

	log.Printf("code: %v", ctx)

    token, err := oAuthConfig.Exchange(ctx, code)
    if err != nil {
        http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
        return // Ensure to return after sending an error response
    }
   
	// Log the JSON data
	log.Printf("Unknown data: %s", jsonData)

	log.Printf("token: %v", token)

    client := oAuthConfig.Client(ctx, token)

	log.Printf("client: %v", token)

    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        http.Error(w, "Failed to get user info", http.StatusInternalServerError)
        return
    }

	log.Printf("Response: %v", resp)

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        http.Error(w, "Failed to get user info", http.StatusInternalServerError)
        return
    }

    var userInfo struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }


    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
        return
    }

	log.Printf("Email: %v", userInfo.Email)
	log.Printf("Name: %v", userInfo.Name)

	http.Redirect(w, request, "http://localhost:3000", http.StatusTemporaryRedirect)
}


func LegacyLoginUser(w http.ResponseWriter, request *http.Request) {
	var user DB.User
	db := DB.DBConnect()

	jsonDecoderError := json.NewDecoder(request.Body).Decode(&user)
	if jsonDecoderError != nil {
		utils.BadResponse(jsonDecoderError, "Error occured during decoding", w)
		return
	}
	
	validate := utils.GetValidator()
	rules := map[string]string{
		"Email":  "required",
		"Password": "required",
	}
	
	validate.RegisterStructValidationMapRules(rules, DB.User{})

	if validationErrors := validate.Struct(user); validationErrors != nil {
		utils.HandleValidationError(validationErrors, w)
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
		utils.BadResponse("Password doesnt match" ,"Error occured",w)
		return
	}
	message := "User successfully logged in!"
	utils.OkResponse(user, message, w)
}