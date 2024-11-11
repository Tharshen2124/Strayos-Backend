package utils

import (
	"errors"
	"example/main/Models"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userData Models.User) (string, error) {
	jwtKey := []byte(GetEnv("JWT_KEY"))
	
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userData.UserId,
			"username": userData.Username,
			"email": userData.Email,
		})
	signedJwtToken, err := jwtToken.SignedString(jwtKey)
	
	// Log the token for debugging purposes
	log.Printf("Generated JWT Token: %s", signedJwtToken)

	return signedJwtToken, err
}

func ParseJWTToken(tokenString string) (string, error) {
    jwtKey := []byte(GetEnv("JWT_KEY"))

    ParsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtKey, nil
    })

    if err != nil {
        log.Printf("Error parsing token: %v", err)
        return "", err
    }

    claims, ok := ParsedToken.Claims.(jwt.MapClaims)
    if !ok || !ParsedToken.Valid {
        return "", errors.New("invalid token or claims")
    }

    log.Printf("Claims: %v", claims)

    userID, exists := claims["user_id"]
    if !exists {
        log.Print("line 45: User ID not found in token")
        return "", errors.New("user ID not found in token")
    }

    // Check the type of user_id
    log.Printf("User ID type: %T, Value: %v", userID, userID)

    var userIDStr string

    // If user_id is a float64 (which happens if it's stored as a number), convert it to string
    switch v := userID.(type) {
    case string:
        userIDStr = v
    case float64:
        userIDStr = fmt.Sprintf("%.0f", v) // Convert float64 to string
    default:
        log.Print("line 45: User ID is of unexpected type")
        return "", errors.New("user ID is of unexpected type")
    }

    return userIDStr, nil
}
