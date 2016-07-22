package main

import (
	// "github.com/carlqt/geodude/geocode"
	// "github.com/iris-contrib/template/html"
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

func main() {
	// apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	// url := "https://maps.googleapis.com/maps/api/geocode/json"

	// g := geocode.GoogleGeoCode{URL: url, ApiKey: apiKey}
	// lng, lat := g.Geocode("xxlkajflkasdjfx")

	// fmt.Printf("Latitude is %f and Longitude is %f", lat, lng)

	iris.StaticServe("./assets")
	iris.Use(logger.New(iris.Logger))
	iris.Post("/search", Search)
	iris.Get("/", Index)

	errorLogger := logger.New(iris.Logger)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Write("My Custom 404 error page ")
	})

	iris.Listen(":8000")
}

func Search(c *iris.Context) {
	u := User{Name: "Madeline", Age: 16}
	c.JSON(iris.StatusOK, u)
}

func Index(c *iris.Context) {
	c.MustRender("hi.html", struct{ Name string }{Name: "iris"})
}
