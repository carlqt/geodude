package main

import (
	"github.com/carlqt/geodude/geocode"
	"github.com/carlqt/geodude/models"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/shopspring/decimal"
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
		api.DELETE("/property/:id", paramToInt(), apiDelete)
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
		p := models.NearbyProperty(point.Geometry.Location["lat"], point.Geometry.Location["lng"])
		c.JSON(http.StatusOK, p)
	}
}

func apiIndex(c *gin.Context) {
	p := models.AllProperties()
	c.JSON(http.StatusOK, p)
}

func apiCreate(c *gin.Context) {
	var err error

	property := &models.Property{Address: c.PostForm("address"),
		Price: strToCurrency(c.PostForm("price")),
		Description: c.PostForm("description"),
		Type: c.PostForm("type"),
		Name: c.PostForm("name"),
	}

	property, err = property.GeocodeAndCreate(g)

	_ = "breakpoint"
	if err != nil {
		color.Red(err.Error())
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, property)
	}
}

func apiGeocode(c *gin.Context) {
	result, err := g.Geocode(c.Query("location"))

	if err != nil {
		color.Red(err.Error())
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"lng": result.Geometry.Location["lng"],
			"lat": result.Geometry.Location["lat"],
		})
	}
}

func apiDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.PropertyDelete(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "deleted",
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

func strToCurrency(str string) decimal.Decimal {
	price, _ := decimal.NewFromString(str)
	return price
}