package DB

type User struct {
	UserId 			string	`json:"UserId"`
	Name 			string	`json:"Name"`
	Email  			string	`json:"Email"`
	Password 		string	`json:"Password"`
	CreatedAt		string	`json:"CreatedAt"`
	UpdatedAt  		string	`json:"UpdatedAt"`
}

type StrayPet struct {
	StrayPetId		string	`json:"PetId"`
	Animal 			string	`json:"Animal"`
	Breed 			string	`json:"Breed"`
	CreatedAt		string	`json:"CreatedAt"`
	UpdatedAt		string	`json:"UpdatedAt"`
}

type MissingPets struct {
	MissingPetsId 	string	`json:"MissingPetsId"`
	PetId			string	`json:"PetId"` 
	LastSeenDate	string	`json:"LastSeenDate"`
	LastSeenTime	string	`json:"LastSeenTime"`
	CreatedAt		string	`json:"CreatedAt"`
	UpdatedAt		string	`json:"UpdatedAt"`
}
