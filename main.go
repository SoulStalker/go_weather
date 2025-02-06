package main

import (
	"demo/weaher/geo"
	"demo/weaher/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(weatherData)
}

