package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	if res.TotalPage%2 != 0 {
		res.TotalPage += 1
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
