package DB

import (
	"time"
)

type User struct {
	UserId 			string	`json:"UserId" gorm:"primaryKey;autoIncrement"`
	Username 		string	`json:"Username" `
	Email			string	`json:"Email"`
	Password 		string	`json:"Password" `
	CreatedAt		time.Time	`json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt  		time.Time	`json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type StrayPet struct {
	StrayPetId		string	`json:"StrayPetId"`
	Animal 			string	`json:"Animal"`
	Breed 			string	`json:"Breed"`
	Latitude		string	`json:"Latitude"`
	Longitude		string	`json:"Longitude"`
	CreatedAt		time.Time	`json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt  		time.Time	`json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type MissingPets struct {
	MissingPetsId 	string	`json:"MissingPetsId"`
	PetId			string	`json:"PetId"` 
	LastSeenDate	string	`json:"LastSeenDate"`
	LastSeenTime	string	`json:"LastSeenTime"`
	CreatedAt		string	`json:"CreatedAt"`
	UpdatedAt		string	`json:"UpdatedAt"`
}
