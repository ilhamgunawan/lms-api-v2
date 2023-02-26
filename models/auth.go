package models

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(user User) (token string, err error) {
	issuedAt := time.Now().Unix()
	expiredAt := time.Now().Add(time.Hour * 2).Unix() // Token expired in 2 hour

	claims := jwt.MapClaims{
		"user_name":  user.Username,
		"user_id":    user.UserID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"iat":        issuedAt,
		"exp":        expiredAt,
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

func VerifyToken(token string) (userId string, err error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return userId, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return userId, err
	}

	userId = claims["user_id"].(string)

	if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok && !parsedToken.Valid {
		return userId, err
	}

	return userId, nil
}
