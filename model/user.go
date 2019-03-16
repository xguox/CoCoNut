package model

import (
	"github.com/xguox/coconut/db"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username       string
	Email          string
	PasswordDigest string
}

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	PasswordDigest, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordDigest = string(PasswordDigest)
	return nil
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordDigest)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindUserByName(username string) (User, error) {
	var user User
	if err := db.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func FindUserByEmail(email string) (User, error) {
	var user User
	if err := db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// UserValidator ...
type UserValidator struct {
	UserTmp struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email,is-uniq"`
		Password string `json:"password" binding:"required,min=6,max=64"`
	} `json:"user"`
	UserModel User
}

func (c *UserValidator) TableName() string {
	return "users"
}

func (uv *UserValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(uv, b)
	if err != nil {
		return err
	}
	uv.UserModel.Username = uv.UserTmp.Username
	uv.UserModel.Email = uv.UserTmp.Email
	uv.UserModel.setPassword(uv.UserTmp.Password)
	return nil
}

// LoginValidator ...
type LoginValidator struct {
	UserTmp struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=64"`
	} `json:"user"`
	UserModel User `json:"-"`
}

func (lv *LoginValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(lv, b)
	if err != nil {
		return err
	}
	lv.UserModel.Email = lv.UserTmp.Email
	return nil
}
