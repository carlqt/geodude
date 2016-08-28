package main

import (
	"fmt"
	"github.com/carlqt/geodude/controllers/properties"
	"github.com/carlqt/geodude/controllers/user"
	"github.com/carlqt/geodude/geocode"
	"github.com/carlqt/geodude/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name  string
	Title string
}

func init() {
	models.InitDB()
	geocode.InitGoogleMap()
}

func main() {
	router := gin.Default()

	// router := gin.New() // Sets gin without default middleware
	// router.Use(gin.Logger())	// Global middleware to add Logger
	// router.Use(beforePong()) // Global middleware to add custom middleware

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", beforePong(), pong)
	router.GET("/newping", properties.Pong)
	router.GET("/", Index)
	router.GET("/token/display", displayToken)
	router.POST("/token/create", createToken)

	api := router.Group("/api")
	{
		api.GET("/search", properties.PropertySearch)
		api.GET("/properties", properties.PropertyIndex)
		api.POST("/property", properties.PropertyCreate)
		api.GET("/geocode", properties.PropertyGeocode)
		api.DELETE("/property/:id", paramToInt(), properties.PropertyDelete)

		api.POST("/user", user.Create)
	}

	router.Run(":8000")
}

func Index(c *gin.Context) {
	u := User{Name: "Iris", Title: "Demo"}
	c.HTML(http.StatusOK, "application.html", u)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// example of custom middleware
func beforePong() gin.HandlerFunc {
	// can add initializers here
	return func(c *gin.Context) {
		c.String(200, "Before Pong ")
		c.Abort()
		//c.Next()
	}
}

func paramToInt() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"error": "bad parameter value",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func createToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   "1",
		"iat":  time.Now().Local(),
		"type": "agent",
	})

	tokenString, err := token.SignedString([]byte("glassdoor"))

	if err != nil {
		c.String(400, err.Error())
	} else {
		c.String(200, tokenString)
	}
}

func displayToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDE2LTA4LTAyVDE4OjAxOjU3LjM0MzE2NDI3KzA4OjAwIiwiaWQiOiIxIiwidHlwZSI6ImFnZW50In0.gRdkq1qNrN3-cyiQEGyUauYlsGlJSTmTKkLUq1K3M7g"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error")
		}

		return []byte("glassdoor"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(200, claims)
	} else {
		c.String(400, err.Error())
	}
}

func jwtAuthenticater() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validateToken(c.Request.Header)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func validateToken(h http.Header) error {
	authHeader := h.Get("Authorization")

	if authHeader == "" {
		return fmt.Errorf("Authorization header not found")
	}

	headerString := strings.Split(authHeader, " ")

	if headerString[0] != "Bearer" {
		return fmt.Errorf("Invalid authorization type")
	}

	_, err := jwt.Parse(headerString[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid algorithm type")
		}

		return []byte("glassdoor"), nil
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}
