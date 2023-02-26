package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgunawan/lms-api-v2/db"
	"github.com/ilhamgunawan/lms-api-v2/models"
)

type GetUsersResponse struct {
	Users        []models.UserAccount `json:"users"`
	Offset       int                  `json:"offset"`
	Limit        int                  `json:"limit"`
	TotalCurrent int                  `json:"total_current"`
	TotalAll     int                  `json:"total_all"`
	TotalPage    int                  `json:"total_page"`
}

func GetUsers(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	users, err := models.GetUsers(int(offset), int(limit))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	count, err := models.CountUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	res := GetUsersResponse{
		Users:        users,
		Offset:       offset,
		Limit:        limit,
		TotalCurrent: len(users),
		TotalAll:     int(count),
	}

	res.TotalPage = res.TotalAll / res.Limit

	if res.TotalAll%res.Limit != 0 {
		res.TotalPage += 1
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"date_of_birth"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var body CreateUserRequest
	err := c.ShouldBindJSON(&body)

	if err != nil || body.BirthDate == "" || body.FirstName == "" || body.LastName == "" || body.Gender == "" || body.Username == "" || body.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing fields"})
		return
	}

	ua := db.UserAccount{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Gender:    body.Gender,
		BirthDate: body.BirthDate,
	}

	ul := db.UserLoginData{
		Username: body.Username,
	}

	user, err := models.CreateUser(ua, ul, body.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserById(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No records found"})
		return
	}

	user, err := models.DeleteUser(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUserById(c *gin.Context) {
	userId := c.Param("id")
	body := db.UserAccount{}

	// x, err := c.GetRawData()

	// y := string(x)

	// if err != nil {

	// }

	// err = json.Unmarshal(x, &body)

	err := c.ShouldBindJSON(&body)

	if err != nil || userId == "" || body.BirthDate == "" || body.FirstName == "" || body.LastName == "" || body.Gender == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing fields"})
		return
	}

	body.ID = userId

	user, err := models.UpdateUser(body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
