package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	res, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("NOT 200")
	}
	body, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}
