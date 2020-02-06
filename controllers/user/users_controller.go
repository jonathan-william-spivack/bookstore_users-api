package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string)(int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil{
		return 0, errors.NewBadRequest("user id should be a number")
	}
	return userId, nil
}


func CreateUser(c *gin.Context){
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil{
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, err := getUserId(c.Param("user_id"))
	if err != nil{
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil{
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context){
	userId, err := getUserId(c.Param("user_id"))
	if err != nil{
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil{
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context){
	userId, IdErr := getUserId(c.Param("user_id"))
	if IdErr != nil{
		c.JSON(IdErr.Status, IdErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil{
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}