package utils

import (
	"errors"
	"example/main/Models"
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
		return "", err
	}

	claims, ok := ParsedToken.Claims.(jwt.MapClaims); 
	
	if ok && ParsedToken.Valid {
		userID, isUserID := claims["user_id"].(string)
		if !isUserID {
			return "", errors.New("user ID not found in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
