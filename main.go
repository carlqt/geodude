package main

import (
	"fmt"
	"github.com/carlqt/geodude/geocode"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	url := "https://maps.googleapis.com/maps/api/geocode/json"

	g := geocode.GoogleGeoCode{URL: url, ApiKey: apiKey}
	lng, lat := g.Geocode("Ubi Avenue 1")

	fmt.Printf("Latitude is %f and Longitude is %f", lat, lng)
}
