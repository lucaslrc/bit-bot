package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type weather []struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type mainJsn struct {
	Temp       float32 `json:"temp"`
	FeelsLike  float32 `json:"feels_like"`
	TempMin    float32 `json:"temp_min"`
	TempMax    float32 `json:"temp_max"`
	Pressure   float32 `json:"pressure"`
	Humidity   float32 `json:"humidity"`
	Visibility float32 `json:"visibility"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
}

type jsonFileWeather struct {
	Weather  weather `json:"weather"`
	MainJsn  mainJsn `json:"main"`
	Wind     wind    `json:"wind"`
	NameCity string  `json:"name"`
}

func getWeather(weatherToken, cityName, stateCode, countryCode string) string {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" +
		cityName + "," +
		stateCode + "," +
		countryCode + "&units=metric&appid=" +
		weatherToken)

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dataWeather jsonFileWeather

	err = json.Unmarshal(body, &dataWeather)

	result := "Temperature in " + dataWeather.NameCity + ": " + fmt.Sprintf("%.1f°C", dataWeather.MainJsn.Temp)

	return result
}
