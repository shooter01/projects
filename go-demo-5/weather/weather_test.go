package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"fmt"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Arange - подготовка данных

	city := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	str := weather.GetWeather(city, 1)
	fmt.Println(str)
	if str != "London" {
		t.Errorf("Ожидалось %v, получение %v", "London", str)
	}

}
