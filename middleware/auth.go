package middleware

import (
	"github.com/xguox/coconut/config"
	"github.com/xguox/coconut/db"
	"github.com/xguox/coconut/model"
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var ErrTokenInvalid = errors.New("invalid token")

var SecretBaseKEY = config.Conf.Server.SecretBaseKEY

type CustomClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// UpdateContextCurrentUser(c, 0)
		token := c.Request.Header.Get("Authorization")
		claims, err := parseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		userID := uint(claims.ID)
		UpdateContextCurrentUser(c, userID)
	}
}

func parseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 与 GenerateToken 的 SignedString对应
		return []byte(SecretBaseKEY), nil
	})
	if err != nil {
		// https://github.com/dgrijalva/jwt-go/blob/master/errors.go
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func GenerateToken(id uint) string {
	claims := CustomClaims{ID: id}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := jwtToken.SignedString([]byte(SecretBaseKEY))
	return token
}

func UpdateContextCurrentUser(c *gin.Context, currentUserID uint) {
	var currentUser model.User

	if err := db.GetDB().First(&currentUser, currentUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.Set("current_user_id", currentUserID)
	c.Set("current_user", currentUser)
}
