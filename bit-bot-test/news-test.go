package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type articles struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type jsonFileNews struct {
	Article articles `json:"articles"`
}

func getNews(tokenNews, countryCode string) []string {
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?country=" + countryCode + "&apiKey=" + tokenNews)

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dataNews jsonFileNews

	err = json.Unmarshal(body, &dataNews)

	arrayResult := make([]string, 2)

	arrayResult[0] = dataNews.Article.Title
	arrayResult[1] = dataNews.Article.URL

	return arrayResult
}
