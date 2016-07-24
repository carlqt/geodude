package main

import (
	"os"
	"github.com/carlqt/geodude/geocode"
	"github.com/carlqt/geodude/models"
	"github.com/gin-gonic/gin"
	"github.com/fatih/color"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name string
	Title string
}

func init() {
	models.InitDB()
}

var g geocode.GoogleGeoCode

func init() {
	g.URL = "https://maps.googleapis.com/maps/api/geocode/json"
	g.ApiKey = os.Getenv("GOOGLE_SERVER_KEY")
}

func main() {
	router := gin.Default()

	// router := gin.New() // Sets gin without default middleware
  // router.Use(gin.Logger())	// Global middleware to add Logger
  // router.Use(beforePong()) // Global middleware to add custom middleware


	router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*")

	router.GET("/ping", beforePong(), pong)
	router.GET("/", Index)

	api := router.Group("/api")
	{
		api.GET("/search", apiSearch)
		api.GET("/properties", apiIndex)
		api.POST("/property", apiCreate)
		api.GET("/geocode", apiGeocode)
	}

	router.Run(":8000")
}

func Index(c *gin.Context) {
	u := User{Name: "Iris", Title: "Demo"}
	c.HTML(http.StatusOK, "application.html", u)
}


func apiSearch(c *gin.Context) {
	point, err := g.Geocode(c.Query("location"))
	if err != nil {
		color.Red(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
			})
	} else {
		p := models.NearbyProperty(point["lat"], point["lng"])	
		c.JSON(http.StatusOK, p)
	}
}

func apiIndex(c *gin.Context) {
	p := models.AllProperties()
	c.JSON(http.StatusOK, p)
}

func apiCreate(c *gin.Context) {
}

func apiGeocode(c *gin.Context) {
	point, err := g.Geocode(c.Query("location"))
	_ = "breakpoint"

	if err != nil {
		color.Red(err.Error())
		c.JSON(500, gin.H{
			"error": err.Error(),
			})
	} else {
		c.JSON(http.StatusOK, gin.H{
				"lng": point["lng"],
				"lat": point["lat"],
			})
	}
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// example of custom middleware
func beforePong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "Before Pong ")
		c.Next()
	}
}