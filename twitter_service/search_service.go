package twitterservice

import (
	"encoding/json"
	"fmt"
	"io"
	"just_news/configuration"
	"just_news/models"
	"net/http"
	"net/url"
	"strings"
)

func Search() (models.ResponseData, error) {

	searchUrl := configuration.SearchUrl()
	query := build_query(configuration.Query())

	encoded_url := url.QueryEscape(query)
	full_url := searchUrl + encoded_url
	client := &http.Client{}

	fmt.Printf("Using url: %s\n", full_url)
	request, err := http.NewRequest(http.MethodGet, full_url, nil)
	if err != nil {
		return models.ResponseData{}, fmt.Errorf("error creating GET request: %s", err)
	}

	request.Header.Add("Authorization", "Bearer "+configuration.Token())

	resp, err := client.Do(request)
	if err != nil {
		return models.ResponseData{}, fmt.Errorf("error sending get request: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ResponseData{}, fmt.Errorf("error reading body: %s", err)
	}

	var tweetData models.ResponseData = models.ResponseData{}
	err = json.Unmarshal(body, &tweetData)
	if err != nil {
		return models.ResponseData{}, fmt.Errorf("error unmarshalling json: %s", err)
	}

	return tweetData, nil
}

func build_query(query models.Query) string {
	required := strings.Join(query.Required, " ")
	optional := strings.Join(query.Optional, " OR ")

	query_string := required + " (" + optional + ")"
	return query_string
}
