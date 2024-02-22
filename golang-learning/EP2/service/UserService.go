package service

import (
	"EP2/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func FindAllUsers(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	ctx.JSON(http.StatusOK, users)
}

// Get User by Id

func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error())
		return
	}

	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "successfully")
}

func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
