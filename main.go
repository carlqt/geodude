package main

import (
	// "github.com/carlqt/geodude/geocode"
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

func main() {
	// apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	// url := "https://maps.googleapis.com/maps/api/geocode/json"

	// g := geocode.GoogleGeoCode{URL: url, ApiKey: apiKey}
	// lng, lat := g.Geocode("xxlkajflkasdjfx")

	// fmt.Printf("Latitude is %f and Longitude is %f", lat, lng)

	// TODO endpoints: Add properties, show all properties in DB, Edit?, Search within radius
	iris.StaticServe("./assets")
	iris.Use(logger.New(iris.Logger))
	iris.Post("/search", Search)
	iris.Get("/", Index)
	iris.Get("/properties", propertyIndex)

	errorLogger := logger.New(iris.Logger)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Write("404 bad page ")
	})

	iris.Listen(":8000")
}

func Search(c *iris.Context) {
	u := User{Name: "Madeline", Age: 16}
	c.JSON(iris.StatusOK, u)
}

func Index(c *iris.Context) {
	u := User{Name: "Maddy", Age: 30}
	c.MustRender("application.html", u)
}

func propertyIndex(c *iris.Context) {
	p := models.AllProperties()
	c.JSON(iris.StatusOK, p)
}
