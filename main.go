package main

import (
	"github.com/carlqt/geodude/geocode"
	"github.com/carlqt/geodude/models"
	"github.com/carlqt/geodude/controllers/properties"
	"github.com/carlqt/geodude/controllers/user"
	// "github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
