package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//USD struct JSON
type USD struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"vardBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type jsonFileMoney struct {
	Usd USD `json:"USD"`
}

func getMoney() []string {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/all/USD-BRL")

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dataMoney jsonFileMoney

	err = json.Unmarshal(body, &dataMoney)

	arrayResult := make([]string, 11)

	arrayResult[0] = dataMoney.Usd.Bid

	return arrayResult
}
