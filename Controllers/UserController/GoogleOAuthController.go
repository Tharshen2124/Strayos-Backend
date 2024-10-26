package UserController

import (
	// "context"
	// "encoding/json"
	// "example/main/DB"
	// "example/main/utils"
	// "log"
	"encoding/json"
	"example/main/DB"
	"example/main/Models"
	"example/main/utils"
	"log"
	"net/http"
	// "os"
	// "github.com/joho/godotenv"
	// "golang.org/x/oauth2"
)

type RequestData struct {
	AccessToken string `json:"accessToken"`
}

type ResponseData struct {
	Sub 			string `json:"sub"`
	Name 			string `json:"name"`
	GivenName 		string `json:"given_name"`
	FamilyName 		string `json:"family_name"`
	Picture 		string `json:"picture"`
	Email 			string `json:"email"`
}


// var oAuthConfig *oauth2.Config

// var userInfo struct {
// 	Email string `json:"email"`
// 	Name  string `json:"name"`
// 	GoogleID string `json:"id"`
// }

// func SignupWithGoogleOAuth(w http.ResponseWriter, request *http.Request) {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	oAuthConfig = &oauth2.Config{
// 		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		Scopes:       []string{"openid", "email", "profile"},
// 		RedirectURL:  "http://localhost:8000/auth/callback",
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
// 			TokenURL: "https://oauth2.googleapis.com/token",
// 		},
// 	}

// 	url := oAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	http.Redirect(w, request, url, http.StatusTemporaryRedirect)
// }

// func CallBack(w http.ResponseWriter, request *http.Request) {
// 	var userData DB.User
// 	db := DB.DBConnect()
// 	ctx := context.Background()
	
// 	code := request.URL.Query().Get("code")
// 	if code == "" {
// 		http.Error(w, "No code in the request", http.StatusBadRequest)
// 		return
// 	}

// 	token, err := oAuthConfig.Exchange(ctx, code)
// 	if err != nil {
// 		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
// 		return
// 	}

// 	client := oAuthConfig.Client(ctx, token)
// 	clientResponse, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
// 		return
// 	}

// 	defer clientResponse.Body.Close()
// 	if clientResponse.StatusCode != http.StatusOK {
// 		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
// 		return
// 	}

// 	log.Printf("Client Response %v", clientResponse)

// 	if err := json.NewDecoder(clientResponse.Body).Decode(&userInfo); err != nil {
// 		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
// 		return
// 	}

// 	if err := db.Where("email = ?", userInfo.Email).First(&userData).Error; err != nil {
// 		log.Printf("Error occured during querying: %v", err.Error())

// 		if err.Error() == "record not found" {
// 			userData.Username = userInfo.Name
// 			userData.Email = userInfo.Email
// 			userData.Provider = "Gmail"
// 			userData.GoogleID = userInfo.GoogleID

// 			result := db.Create(&userData)
// 			if result.Error != nil {
// 				log.Printf("Error inserting data: %v", result.Error)
// 			} else {
// 				log.Printf("Successfully inserted data!")
// 			}
// 		}
// 	} else {
// 		if userData.GoogleID == "" {
// 			userData.GoogleID = userInfo.GoogleID
// 			userData.Provider = "Gmail"

// 			db.Save(&userData)

// 			signedJwtToken, jwtTokenError := utils.CreateToken(userData)
// 			if jwtTokenError != nil {
// 				log.Printf("Error in creating JWT Token: %v", jwtTokenError)
// 			}

// 		} else {
// 			signedJwtToken, jwtTokenError := utils.CreateToken(userData)
// 			if jwtTokenError != nil {
// 				log.Printf("Error in creating JWT Token: %v", jwtTokenError)
// 			}
// 			http.Redirect(w, request, "http://localhost:3000/signup", http.StatusTemporaryRedirect)
// 		}
// 	}
// }

func SignupWithGoogleOAuth(w http.ResponseWriter, request *http.Request) {
	var accessToken RequestData
	var userInfo ResponseData
	var userData Models.User
	db := DB.DBConnect()
	// log.Print(request.Body)
	jsonDecoderError := json.NewDecoder(request.Body).Decode(&accessToken)

	if jsonDecoderError != nil {
		log.Printf("JSON Decoder error: %v",  jsonDecoderError)
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken.AccessToken)

	if err != nil {
		log.Printf("Error during fetching: %v", err)
	} else {
		jsonDecoderError := json.NewDecoder(response.Body).Decode(&userInfo)
		if jsonDecoderError != nil {
			log.Printf("Error during decoding: %v", jsonDecoderError)
		}
	}

	if err := db.Where("email = ?", userInfo.Email).First(&userData).Error; err != nil {
		log.Printf("Error occured during querying: %v", err.Error())

		if err.Error() == "record not found" {
			userData.Username = userInfo.Name
			userData.Email = userInfo.Email
			userData.Provider = "Gmail"
			userData.GoogleID = userInfo.Sub

			result := db.Create(&userData)
			if result.Error != nil {
				log.Printf("Error inserting data: %v", result.Error)
			} else {
				log.Printf("Successfully inserted data!")

				signedJwtToken, jwtTokenError := utils.CreateToken(userData)
				if jwtTokenError != nil {
					log.Printf("Error in creating JWT Token: %v", jwtTokenError)
					utils.BadResponse(jwtTokenError, w)
				}

				utils.AuthOkResponse(signedJwtToken, w)
			}
		}
	} else {
		if userData.GoogleID == "" {
			userData.GoogleID = userInfo.Sub
			userData.Provider = "Gmail"

			db.Save(&userData)

			signedJwtToken, jwtTokenError := utils.CreateToken(userData)
			if jwtTokenError != nil {
				log.Printf("Error in creating JWT Token: %v", jwtTokenError)
				utils.BadResponse(jwtTokenError, w)
			}

			utils.AuthOkResponse(signedJwtToken, w)
		} else {
			signedJwtToken, jwtTokenError := utils.CreateToken(userData)
			if jwtTokenError != nil {
				log.Printf("Error in creating JWT Token: %v", jwtTokenError)
				utils.BadResponse(jwtTokenError, w)
			}

			utils.AuthOkResponse(signedJwtToken, w)
		}
	}
}
