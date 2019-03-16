package controller

import (
	"github.com/xguox/coconut/middleware"
	"github.com/xguox/coconut/model"
	"github.com/xguox/coconut/util"
	"github.com/xguox/coconut/serializer"
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
		if user.CheckPassword(password) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "密码错误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"email": user.Email})
		}
	}
}

func UserLogin(c *gin.Context) {
	v := model.LoginValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}
	user, err := model.FindUserByEmail(v.UserModel.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not Registered email or invalid password"})
		return
	}

	if user.CheckPassword(v.UserTmp.Password) != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid password"})
		return
	}
	middleware.UpdateContextCurrentUser(c, user.ID)
	serializer := serializer.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
