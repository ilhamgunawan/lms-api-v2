package models

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(user User) (token string, err error) {
	claims := jwt.MapClaims{
		"user_name":  user.Username,
		"user_id":    user.UserID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"exp":        time.Now().Add(time.Hour),
	}

	// Create access token
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Sign token
	token, err = jwtClaims.SignedString(jwtSecret)

	if err != nil {
		return token, err
	}

	return token, nil
}
