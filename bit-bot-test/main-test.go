package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type listIDGroup struct {
	Groups string `json:"group-id"`
}

type listIDChannel struct {
	Channels string `json:"channel-id"`
}

type config struct {
	TelegramBOTToken    string `json:"telegram-bot-token"`
	OpenWeatherMAPToken string `json:"open-weather-map-token"`
	NewsAPIToken        string `json:"news-api-token"`
	CityName            string `json:"city-name"`
	StateCode           string `json:"state-code"`
	CountryCode         string `json:"country-code"`
	MoneyConvert        string `json:"money-convert"`
}

type jsonFile struct {
	Config       config        `json:"config"`
	ListGroups   listIDGroup   `json:"groups"`
	ListChannels listIDChannel `json:"channels"`
}

//JSONConfig is a configuration file
var JSONConfig jsonFile

func main() {
	fmt.Println("initiating bot...")
	time.Sleep(2 * time.Second)
	initBOT()
}

func initBOT() {
	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create("config.json")
			if err != nil {
				log.Fatal(err)
				return
			}

			var jsn jsonFile

			byteJSON, err := json.MarshalIndent(jsn, "	", "  ")

			if err != nil {
				log.Fatal(err)
				return
			}

			file.Write(byteJSON)
			log.Println("Configure the new .json file created into this folder.")
			return

		}
		log.Fatal(err)
		return
	}

	defer file.Close()

	byteValues, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValues, &JSONConfig)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(getWeather(JSONConfig.Config.OpenWeatherMAPToken, JSONConfig.Config.CityName, JSONConfig.Config.StateCode, JSONConfig.Config.CountryCode))
	// initTelegramService(JSONConfig.Config.TelegramBOTToken)
}
