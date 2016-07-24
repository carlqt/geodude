package main

import (
	"os"
	"github.com/carlqt/geodude/geocode"
	// "github.com/iris-contrib/template/html"
	"github.com/carlqt/geodude/models"
	// "github.com/iris-contrib/middleware/logger"
	// "github.com/kataras/iris"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name string
	Age  int
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

	router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*")

	router.GET("/ping", pong)
	router.GET("/search", Search)
	router.GET("/", Index)
	router.GET("/properties", propertyIndex)
	router.POST("/property", propertyCreate)
	router.GET("/convert", convert)

	router.Run(":8000")
}

func Search(c *gin.Context) {
	point, err := g.Geocode(c.Query("location"))
	if err != nil {
		// c.EmitError(iris.StatusInternalServerError)
	} else {
		p := models.NearbyProperty(point["lat"], point["lng"])	
		c.JSON(http.StatusOK, p)
	}
}

func Index(c *gin.Context) {
	u := User{Name: "Iris", Age: 30}
	c.HTML(http.StatusOK, "application.html", u)
}

func propertyIndex(c *gin.Context) {
	p := models.AllProperties()
	c.JSON(http.StatusOK, p)
}

func propertyCreate(c *gin.Context) {
}

func convert(c *gin.Context) {
	lat, lng := g.Geocode(c.Query("location"))

	c.JSON(http.StatusOK, gin.H{
			"lng": lng,
			"lat": lat,
		})
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}