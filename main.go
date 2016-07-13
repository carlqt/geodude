package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var address string

	apiKey := os.Getenv("GOOGLE_SERVER_KEY")
	url := "https://maps.googleapis.com/maps/api/geocode/json"

	fmt.Println("Enter an address: ")
	_, err := fmt.Scanln(&address)
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

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
