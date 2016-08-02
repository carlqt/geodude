package main

import (
	"fmt"
	"github.com/carlqt/geodude/controllers"
	"github.com/carlqt/geodude/geocode"
	"github.com/carlqt/geodude/models"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	router.GET("/newping", controllers.Pong)
	router.GET("/", Index)
	router.GET("/token/display", displayToken)
	router.POST("/token/create", createToken)

	api := router.Group("/api")
	{
		api.GET("/search", controllers.PropertySearch)
		api.GET("/properties", jwtAuthenticate(), controllers.PropertyIndex)
		api.POST("/property", controllers.PropertyCreate)
		api.GET("/geocode", controllers.PropertyGeocode)
		api.DELETE("/property/:id", paramToInt(), controllers.PropertyDelete)
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

func jwtAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header["Authorization"])
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
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDE2LTA4LTAyVDE4OjAxOjU3LjM0MzE2NDI3KzA4OjAwIiwiaWQiOiIxIiwidHlwZSI6ImFnZW50In0.gRdkq1qNrN3-cyiQEGyUauYlsGlJSTmTKkLUq1K3M7g"
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
