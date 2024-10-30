package StrayPetsController

import (
	"encoding/json"
	"example/main/DB"
	"example/main/Models"
	"example/main/SDKs"
	"example/main/utils"
	"log"
	"net/http"
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

func Index(w http.ResponseWriter, request *http.Request) {
    var strayPets []Models.StrayPet
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
}

func Create(w http.ResponseWriter, request *http.Request) {
    var strayPet Models.StrayPet
    var savedStrayPet Models.StrayPet
    db := DB.DBConnect()
    
    cloudinaryInstance, ctx := SDKs.Credentials()
    request.ParseMultipartForm(10 << 20)
    file, handler, err := request.FormFile("Image")

    if(err != nil) {
        utils.BadResponse(err, w)
        return
    }

    strayPet.Animal = request.FormValue("Animal")
    strayPet.Status = request.FormValue("Status")
    strayPet.UserId = request.FormValue("UserId")
    strayPet.Latitude = request.FormValue("Latitude")
    strayPet.Longitude = request.FormValue("Longitude")

    validate := utils.GetValidator()
    rules := map[string]string{
        "Animal": "required",
        "UserId": "required",
        "Status": "required",
        "Latitude": "required",
        "Longitude": "required",
    }

    validate.RegisterStructValidationMapRules(rules, Models.StrayPet{})
    if validationErrors := validate.Struct(strayPet); validationErrors != nil {
        utils.HandleValidationError(validationErrors, w)
        return
	}

    imagePublicID := SDKs.GeneratePrefixedUniqueID()
    SDKs.UploadImage(cloudinaryInstance, ctx, file, imagePublicID)
    ImageURL := SDKs.GetTransformedImage(cloudinaryInstance, ctx, imagePublicID)

    strayPet.ImageURL = ImageURL
    strayPet.ImagePublicID = imagePublicID
    strayPet.Image = handler.Filename

    log.Print(strayPet)

    db.Create(&strayPet)

    db.Preload("User").First(&savedStrayPet, strayPet.StrayPetId)
 
    response := Response{
        Data: savedStrayPet,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
