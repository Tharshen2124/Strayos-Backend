package Models

import (
	"time"
)
type User struct {
	UserId      int       `json:"UserId" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"Username"`
	Email       string    `json:"Email"`
	Password    string    `json:"Password,omitempty"`
	GoogleID    string    `json:"GoogleId,omitempty"`
	Provider    string    `json:"Provider" gorm:"default:'email'"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type StrayPet struct {
	StrayPetId    int       `json:"StrayPetId" gorm:"primaryKey;autoIncrement"`
	UserId        int       `json:"UserId"` // Match the type to int
	Animal        string    `json:"Animal"`
	Status        string    `json:"Status"`
	Latitude      string    `json:"Latitude"`
	Longitude     string    `json:"Longitude"`
	Image         string    `json:"Image,omitempty"`
	ImagePublicID string    `json:"ImagePublicID,omitempty"`
	ImageURL      string    `json:"ImageURL,omitempty"`
	CreatedAt     time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"UpdatedAt" gorm:"autoUpdateTime"`
	User          User      `gorm:"foreignKey:UserId;references:UserId"`
}

// type MissingPets struct {
// 	MissingPetsId 	string	`json:"MissingPetsId"`
// 	PetId			string	`json:"PetId"` 
// 	LastSeenDate	string	`json:"LastSeenDate"`
// 	LastSeenTime	string	`json:"LastSeenTime"`
// 	CreatedAt		string	`json:"CreatedAt"`
// 	UpdatedAt		string	`json:"UpdatedAt"`
// }

