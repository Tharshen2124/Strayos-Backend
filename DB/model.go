package DB

import (
	"time"
)

type User struct {
	UserId 			string		`json:"UserId" gorm:"primaryKey;autoIncrement"`
	Username 		string		`json:"Username" `
	Email			string		`json:"Email"`
	Password 		string		`json:"Password,omitempty"`
	GoogleID  		string    	`json:"GoogleID,omitempty"`           // Store Google ID for users who login via Google
	Provider  		string    	`json:"Provider" gorm:"default:'email'"` // Indicates if user signed up with email or Google
	CreatedAt		time.Time	`json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt  		time.Time	`json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type StrayPet struct {
	StrayPetId		string		`json:"StrayPetId" gorm:"primaryKey;autoIncrement"`
	Animal 			string		`json:"Animal"`
	Status			string		`json:"Status"`
	Latitude		string		`json:"Latitude"`
	Longitude		string		`json:"Longitude"`
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

