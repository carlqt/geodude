package geocode

import "os"

var URL string
var KEY string

func InitGoogleMap() {
  URL = "https://maps.googleapis.com/maps/api/geocode/json"
  KEY = os.Getenv("GOOGLE_SERVER_KEY")
}