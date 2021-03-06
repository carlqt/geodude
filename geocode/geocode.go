package geocode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type JsonResponse struct {
	Status       string       `json:"status"`
	Results      []ResultBody `json:"results"`
	ErrorMessage string       `json:"error_message"`
}

type ResultBody struct {
	FormattedAddress string       `json:"formatted_address"`
	Geometry         GeometryBody `json:"geometry"`
}

type GeometryBody struct {
	Location map[string]float32
}

// Send an HTTP request to https://maps.googleapis.com/maps/api/geocode/json
// Returns an error if can't connect
func request(address string) (geometry *ResultBody, err error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", KEY)
	q.Add("address", address)
	q.Add("components", "country:SG")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var jsn JsonResponse

	dec := json.NewDecoder(res.Body)
	dec.Decode(&jsn)

	if jsn.Status == "OK" {
		geometry = &jsn.Results[0]
		return geometry, nil
	} else {
		return nil, fmt.Errorf("Status: %s\nError Message: %s", jsn.Status, jsn.ErrorMessage)
	}
}

func Geocode(address string) (result *ResultBody, err error) {
	result, err = request(address)

	return result, err
}
