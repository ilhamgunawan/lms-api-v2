package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Username  string `json:"user_name"`
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"date_of_birth"`
	Token     string `json:"token"`
}

func Login(c *gin.Context) {
	var body LoginRequest

	err := c.ShouldBindJSON(&body)

	if err != nil || body.Username == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing fields"})
		return
	}

	user, err := models.GetUserByUsername(body.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password not match"})
		return
	}

	plainPassword := []byte(body.Password)
	hashedPassword := []byte(user.PasswordHash)

	err = bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password not match"})
		return
	}

	token, err := models.CreateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	res := LoginResponse{
		Username:  user.Username,
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		Token:     token,
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

type ValidateTokenBody struct {
	Token string `json:"token"`
}

func ValidateToken(c *gin.Context) {
	var body ValidateTokenBody

	err := c.ShouldBindJSON(&body)

	if err != nil || body.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	err = models.VerifyToken(body.Token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": body})
}
