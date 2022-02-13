package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var searchUrl string = "https://api.twitter.com/2/tweets/search/recent?max_results=100&tweet.fields=entities&query="
var query string = "news (nc OR unitedstates OR usa OR world OR science OR math OR space OR tech OR technology) lang:en -is:retweet is:verified"

type tweet_url struct {
	ExpandedUrl string `json:"expanded_url"`
}

type entities struct {
	Urls []tweet_url `json:"urls"`
}

type tweet struct {
	Entities entities `json:"entities"`
	Lang     string   `json:"lang"`
	Text     string   `json:"text"`
	Id       string   `json:"id"`
}

type responseData struct {
	Tweets []tweet `json:"data"`
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config yaml file: %w \n", err))
	}

	encoded_url := url.QueryEscape(query)
	full_url := searchUrl + encoded_url
	client := &http.Client{}

	fmt.Printf("Using url: %s\n", full_url)
	request, err := http.NewRequest(http.MethodGet, full_url, nil)
	if err != nil {
		log.Fatalf("error creating GET request: %s", err)
	}

	request.Header.Add("Authorization", "Bearer "+viper.GetString("token"))

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("error sending get request: %s", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %s", err)
	}

	var tweetData responseData = responseData{}
	err = json.Unmarshal(body, &tweetData)
	if err != nil {
		log.Fatalf("error unmarshalling json: %s", err)
	}

	file, err := os.Create("./datafile.json")
	if err != nil {
		log.Fatalf("unable to create datafile.json: %s", err)
	}

	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	err = enc.Encode(tweetData)
	if err != nil {
		log.Fatalf("error encoding tweet data to bytes: ", err)
	}
	file.Write(buf.Bytes())
	fmt.Println("Wrote tweet data to file.")

	unique_tweets := map[string]bool{}

	for _, tweet := range tweetData.Tweets {

		formatted_tweet_text := strings.Split(tweet.Text, "https://t.co")[0]

		if !unique_tweets[formatted_tweet_text] {
			unique_tweets[formatted_tweet_text] = true
		} else {
			continue
		}

		fmt.Printf("======================\n%s\n", formatted_tweet_text)

		urls := []string{}

		for _, url := range tweet.Entities.Urls {
			urls = append(urls, url.ExpandedUrl)
		}

		if len(urls) > 0 {
			fmt.Printf("\n\tUrls: %s\n", strings.Join(urls, ", "))
		}

		fmt.Printf("======================\n")
	}

}
