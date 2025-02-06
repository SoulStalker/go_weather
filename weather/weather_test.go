package weather_test

import (
	"demo/weaher/geo"
	"demo/weaher/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Ufa"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	got, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("GetWeather returned: %v", err)
	}
	if !strings.Contains(got, expected) {
		t.Errorf("Expected %v, got %v\n", expected, got)
	}
}

var testCases = []struct {
	name string
	format int
}{
	{ name: "Big fmt", format: 147 }, 
	{ name: "0 fmt", format: 0 },
	{ name: "Mimus fmt", format: -1 },
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Ufa"
			geoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Expected %v, got %v\n", weather.ErrWrongFormat, err)
			}
		})
	}
	
}
