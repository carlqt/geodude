package user

import (
	// "encoding/json"
	"net/http"
	"time"

	"github.com/carlqt/geodude/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create endpoint for user resource
func Create(c *gin.Context) {
	user := &models.User{
		Username:             c.PostForm("username"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("passwordConfirmation"),
		Role:                 c.PostForm("type"),
	}

	if err := user.Validate(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if err := user.Create(); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			webToken := createToken(user)

			cookie := &http.Cookie{
				Name:  "jwt",
				Value: webToken,
			}
			http.SetCookie(c.Writer, cookie)
			c.JSON(200, user)
		}
	}
}

func createToken(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"iat":      time.Now().Local(),
		"type":     user.Role,
		"username": user.Username,
		"email":    user.Email,
	})

	tokenString, _ := token.SignedString([]byte("glassdoor"))
	return tokenString
}
