package UserController

import (
	"net/http"
	"encoding/json"
	"example/main/DB"
	"log"
	"github.com/gorilla/mux"
	"fmt"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// var user DB.User
	value := r.Body
	fmt.Println(value)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user DB.User
	vars := mux.Vars(r)
	userId := vars["userId"]
	db := DB.DBConnect()
	rows , err := db.Query("SELECT * FROM users where user_id = $1", userId)

	if err != nil {
		log.Fatalf("Error during querying : %v", err)
		return 
	} else {
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Fatalf("Error: %v", err)
				return
			} 
		}
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []DB.User
	db := DB.DBConnect()
	rows, err := db.Query("SELECT * FROM users;")
	
	if err != nil {
		log.Fatalf("Error during querying : %v", err)
		return 
	} else {
		defer rows.Close()
		for rows.Next() {
			var user DB.User
			if err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Fatalf("Error: %v", err)
				return
			} 
			users = append(users, user)
		}
	}

	json.NewEncoder(w).Encode(users)
}