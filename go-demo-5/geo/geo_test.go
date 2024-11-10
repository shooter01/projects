package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - подготовка данных
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	// Arange - подготовка данных
	city := "Londonasds"

	// Act - выполняем функцию
	_, err := geo.GetMyLocation(city)

	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrNoCity, err)
	}
}
