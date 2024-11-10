package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"errors"
	"fmt"
	"testing"
)

var testcases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "Format", format: 6},
	{name: "Negative", format: -1},
}

func TestGetWeather(t *testing.T) {
	// Arange - подготовка данных

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			city := geo.GeoData{
				City: "London",
			}

			// Act - выполняем функцию
			str, err := weather.GetWeather(city, tc.format)
			fmt.Println(str)
			if err != errors.New(weather.WRONG_FROMAT) {
				t.Errorf("Ожидалось %v, получение %v", errors.New(weather.WRONG_FROMAT), err)
			}
		})
	}

}
