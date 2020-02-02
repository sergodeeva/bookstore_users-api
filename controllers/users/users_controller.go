package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sergodeeva/bookstore_users-api/domain/users"
	"github.com/sergodeeva/bookstore_users-api/services"
	"github.com/sergodeeva/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	result, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
