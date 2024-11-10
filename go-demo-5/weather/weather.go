package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var WRONG_FROMAT = "WRONG_FORMAT"

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 5 {
		return "", errors.New(WRONG_FROMAT)
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	// if err != nil {
	// 	return ""
	// }

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}
