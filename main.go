package main

import (
	"os"
	"github.com/carlqt/geodude/geocode"
	// "github.com/iris-contrib/template/html"
	"github.com/carlqt/geodude/models"
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
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
	// apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	// url := "https://maps.googleapis.com/maps/api/geocode/json"

	// g := geocode.GoogleGeoCode{URL: url, ApiKey: apiKey}
	// lng, lat := g.Geocode("xxlkajflkasdjfx")

	// fmt.Printf("Latitude is %f and Longitude is %f", lat, lng)

	// TODO endpoints: Add properties, Edit?, Search within radius
	iris.StaticServe("./assets")
	iris.Use(logger.New(iris.Logger))
	iris.Get("/ping", pong)
	iris.Get("/search", Search)
	iris.Get("/", Index)
	iris.Get("/properties", propertyIndex)
	iris.Post("/property", propertyCreate)
	iris.Get("/convert", convert)

	errorLogger := logger.New(iris.Logger)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Write("404 bad page ")
	})

	iris.Listen(":8000")
}

func Search(c *iris.Context) {
	point, err := g.Geocode(c.URLParam("location"))
	if err != nil {
		c.EmitError(iris.StatusInternalServerError)
	} else {
		p := models.NearbyProperty(point["lat"], point["lng"])	
		c.JSON(iris.StatusOK, p)
	}
}

func Index(c *iris.Context) {
	u := User{Name: "Iris", Age: 30}
	c.MustRender("application.html", u)
}

func propertyIndex(c *iris.Context) {
	p := models.AllProperties()
	c.JSON(iris.StatusOK, p)
}

func propertyCreate(c *iris.Context) {
}

func convert(c *iris.Context) {
	lat, lng := g.Geocode(c.URLParam("location"))

	c.JSON(iris.StatusOK, iris.Map{
			"lng": lng,
			"lat": lat,
		})
}

func pong(c *iris.Context) {
	c.Write("pong")
}