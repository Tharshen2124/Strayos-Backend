package utils

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(validationErrors any, w http.ResponseWriter) {
	errorMap := make(map[string]interface{})
    for _, validationError := range validationErrors.(validator.ValidationErrors) {
		validationErrorValue := fmt.Sprintf("This Field with validation '%s' has failed",validationError.ActualTag())
		errorMap[validationError.Field()] = validationErrorValue 
    }
	message := "Error occured during validation"
	BadResponse(errorMap, message, w)
}