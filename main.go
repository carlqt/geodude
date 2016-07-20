package main

import (
	// "github.com/carlqt/geodude/geocode"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name string
	Age int
}

func main() {
	// apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	// url := "https://maps.googleapis.com/maps/api/geocode/json"

	// g := geocode.GoogleGeoCode{URL: url, ApiKey: apiKey}
	// lng, lat := g.Geocode("xxlkajflkasdjfx")

	// fmt.Printf("Latitude is %f and Longitude is %f", lat, lng)

	iris.Use(logger.New(iris.Logger))
	iris.Post("/search", Search)

	errorLogger := logger.New(iris.Logger)


    iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
        errorLogger.Serve(ctx)
        ctx.Write("My Custom 404 error page ")
    })

	iris.Listen(":8080")
}

func Search(c *iris.Context) {
	u := User{Name: "Madeline", Age: 16}
	c.JSON(iris.StatusOK, u)
}