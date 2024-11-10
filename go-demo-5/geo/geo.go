package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPupulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := CheckCity(city)
		if !isCity {
			panic("Такого города нет")
		}
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

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://ipapi.co/json/", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	var populationResponse CityPupulationResponse
	json.Unmarshal(body, &populationResponse)
	return !populationResponse.Error
}
