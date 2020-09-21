package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type googleAPIResponse struct {
	Results Results `json:"results"`
}

// Results of google api
type Results []Geometry

// Geometry object
type Geometry struct {
	Geometry Location `json:"geometry"`
}

// Location object
type Location struct {
	Location Coordinates `json:"location"`
}

// Coordinates object
type Coordinates struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

func fetchLatLong(postalCode string, token string) (float32, float32) {
	url := "https://maps.googleapis.com/maps/api/geocode/json?region=BR&address=" + postalCode + "&key=" + token

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return 0, 0
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0
	}

	var data googleAPIResponse
	json.Unmarshal(body, &data)
	return data.Results[0].Geometry.Location.Latitude, data.Results[0].Geometry.Location.Longitude
}
