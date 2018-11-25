package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type source struct {
	ID string `json:"id"`
}
type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type topicsAPI struct {
	Articles []Topic `json:"articles"`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))
	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)
	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic {
	var topics []Topic
	for _, source := range sources {
		body := getData(topicURL(source.ID))
		var topicAPI topicsAPI
		json.Unmarshal(body, &topicAPI)
		topics = append(topics, topicAPI.Articles...)
	}
	return topics
}

/*
Category:
general
technology
sports
business
entertainment
science

*/
func sourceURL(category string) string {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func topicURL(id string) string {
	return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=ccf6d2edaede4bb5ba860061a06420e3", id)
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
