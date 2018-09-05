package controller

import (
	"coconut/model"
	"coconut/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	v := model.UserValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}

	if err := model.SaveData(&v.UserModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   v.UserModel.ID,
		"name": v.UserModel.Username,
	})
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := model.FindUserByName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no user found"})
	} else {
		if err := user.CheckPassword(password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "密码错误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"email": user.Email})
		}
	}
}
