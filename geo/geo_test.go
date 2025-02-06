package geo_test

import (
	"demo/weaher/geo"
	"testing"
)

// func TestGetMyLocation(t *testing.T) {
// 	// Arrange - подготовка, expected результат
// 	city := "Ufa"
// 	expected := geo.GeoData{
// 		City: "Ufa",
// 	}
// 	// Act - выполняем функцию
// 	got, err := geo.GetMyLocation(city)
// 	// Assert - проверка результата с expected
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if got.City != expected.City {
// 		t.Errorf("Expected %v, got %v\n", expected, got )
// 	}

// }

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londonnasdfs"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Exptected %v, got %v\n", geo.ErrNoCity, err)
	}
}


