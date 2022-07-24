package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"jwt_auth_golang/models"
	u "jwt_auth_golang/utils"
)

// Get one user by id
func GetUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	user := models.GetUser(id)
	if user == nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = user
	u.Respond( c, http.StatusOK, resp)
	return
}

// Get all the users in the users table
func GetUsers(c *gin.Context) {
	resp := u.Message(true, "success")
	users := models.GetUsers()
	if users == nil {
		c.JSON(http.StatusOK, "No users found")
		return
	}
	resp["data"] = users
	u.Respond(c , http.StatusOK, resp)
	return
}

func CreateUser(c *gin.Context) {

	user := &models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	resp := user.Create()
	u.Respond(c, http.StatusCreated, resp)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	err = models.GetUserForUpdateOrDelete(id, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}


	user.ID = uint(id)
	user.UpdatedAt = time.Now().Local()

	// Update user here
	err = models.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Could not update the record")
		return
	}
	resp := u.Message(true, "Updated successfully")
	resp["data"] = user
	u.Respond(c, http.StatusCreated, resp)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "There was an error in your request")
		return
	}

	err = models.GetUserForUpdateOrDelete(id, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	err = models.DeleteUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Could not delete the record")
		return
	}
	u.Respond(c, http.StatusOK, u.Message(true, "User has been deleted successfully"))
}
