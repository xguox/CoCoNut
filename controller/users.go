package controller

import (
	"github.com/xguox/coconut/middleware"
	"github.com/xguox/coconut/model"
	"github.com/xguox/coconut/serializer"
	"github.com/xguox/coconut/util"
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

// @Summary 后台账号登录
// @Description 后台账号登录
// @Accept json
// @Produce json
// @Tags auth
// @Param body body model.LoginValidator true "账号登录请求参数"
// @Success 200 {string} json "{msg:"请求处理成功"}"
// @Failure 401 {string} json "{msg:"账号或密码有误"}"
// @Failure 422 {string} json "{msg:"请求参数有误"}"
// @Failure 500 {string} json "{msg:"服务器错误"}"
// @Router /admin/login [post]
func UserLogin(c *gin.Context) {
	v := model.LoginValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}
	user, err := model.FindUserByEmail(v.UserModel.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Registered email or invalid password"})
		return
	}

	if user.CheckPassword(v.UserTmp.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}
	middleware.UpdateContextCurrentUser(c, user.ID)
	serializer := serializer.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
