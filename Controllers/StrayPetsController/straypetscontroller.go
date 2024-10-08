package StrayPetsController

import (
	"encoding/json"
	"example/main/DB"
	"example/main/utils"
	"fmt"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
	// "github.com/gorilla/websocket"
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

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func Index(w http.ResponseWriter, request *http.Request) {
    var strayPets []DB.StrayPet
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

    // data, marshalError := json.Marshal(strayPets)
    // if marshalError != nil {
    //     fmt.Println("Marshal Error: ", marshalError)
    // }

	// conn, upgradeError := upgrader.Upgrade(w, request, nil)
	// if upgradeError != nil {
	// 	fmt.Println("Error upgrading: ", upgradeError)
	// 	return
	// }

	// defer conn.Close()

	// for {
    //     WriteError := conn.WriteMessage(websocket.TextMessage, data)
	// 	if  WriteError != nil {
	// 		fmt.Println("Error writing message:", WriteError)
	// 		break
	// 	}
	// }
}

func Create(w http.ResponseWriter, request *http.Request) {
    var strayPet DB.StrayPet
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
        "Status": "required",
    }

    validate.RegisterStructValidationMapRules(rules, DB.StrayPet{})
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

    response := Response{
        Data: strayPet,
    }
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// _, message, ReadError := conn.ReadMessage()
// if ReadError != nil {
//     fmt.Println("Error reading message:", ReadError)
//     break
// }
// fmt.Printf("Received: %s \n", message)



// func Index(w http.ResponseWriter, request *http.Request) {
//     conn, err := upgrader.Upgrade(w, request, nil)
//     if err != nil {
//        fmt.Println("Error upgrading:", err)
//        return
//     }
//     defer conn.Close()

//     go handleConnection(conn)
// }

// func handleConnection(conn *websocket.Conn) {
//     // WebSocket Ping/Pong to keep connection alive
//     conn.SetReadDeadline(time.Now().Add(60 * time.Second))
//     conn.SetPongHandler(func(string) error {
//         conn.SetReadDeadline(time.Now().Add(60 * time.Second))
//         return nil
//     })

//     for {
//         _, message, readError := conn.ReadMessage()
//         if readError != nil {
//             if websocket.IsUnexpectedCloseError(readError, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
//                 fmt.Printf("Unexpected close error: %v\n", readError)
//             } else {
//                 fmt.Printf("Connection closed: %v\n", readError)
//             }
//             break
//         }

//         fmt.Printf("Received: %s\n", message)
//         writeError := conn.WriteMessage(websocket.TextMessage, message)

//         if writeError != nil {
//             fmt.Println("Error writing message:", writeError)
//             break
//         }
//     }
// }
