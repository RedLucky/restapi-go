package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	// ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *Connection) LoginHandler(c *gin.Context) {
	var (
		user  User
		error bool
	)

	if c.PostForm("username") != "test" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username",
		})
		error = true
	} else {
		if c.PostForm("password") != "test" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong password",
			})
			error = true
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		error = true
	}
	if !error {
		payload := sign.Claims.(jwt.MapClaims)
		payload["user"] = user.Username
		// payload["id"] = user.ID

		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"payload": payload,
		})
	}
}
