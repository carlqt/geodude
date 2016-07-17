package main

import (
  "encoding/json"
  "bufio"
  "fmt"
  // "io/ioutil"
  // "github.com/d4l3k/go-pry/pry"
  "net/http"
  "os"
  "io"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

type JsonResponse struct {
	Status string `json:"status"`
	Results []ResultBody `json:"results"`
}

type ResultBody struct {
	FormattedAddress string `json:"formatted_address"`
	Geometry GeometryBody `json:"geometry"`
}

type GeometryBody struct {
	Location map[string]float32
}

func main() {
  var address string

  apiKey := os.Getenv("GOOGLE_SERVER_KEY")
  url := "https://maps.googleapis.com/maps/api/geocode/json"

  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Enter an address: ")
  address, err := reader.ReadString('\n')
  checkErr(err)

  request, err := http.NewRequest("GET", url, nil)
  checkErr(err)

  q := request.URL.Query()
  q.Add("key", apiKey)
  q.Add("address", address)
  request.URL.RawQuery = q.Encode()

  client := &http.Client{}

  res, err := client.Do(request)
  checkErr(err)
  defer res.Body.Close()
  displaySomething(res.Body)

  // body, _ := ioutil.ReadAll(res.Body)
  // fmt.Println(string(body))
}


func displaySomething(body io.ReadCloser) {
  var data JsonResponse
  dec := json.NewDecoder(body)
  dec.Decode(&data)

  // fmt.Println(data.Results[0]["geometry"]["location"])
  fmt.Println(data.Results[0].Geometry.Location["lat"])
  fmt.Println(data.Results[0].Geometry.Location["lng"])
  fmt.Println(data.Status)
}
