package main

import (
	"fmt"
	"just_news/configuration"
	"just_news/models"
	twitterservice "just_news/twitter_service"
	"log"
	"strings"
)

func main() {

	err := configuration.Init()
	if err != nil {
		log.Fatalf("error initializing configuration: %s", err)
	}

	tweet_data, err := twitterservice.Search()
	if err != nil {
		log.Fatal(err)
	}

	print_tweets(tweet_data)
}

func print_tweets(tweet_data models.ResponseData) {

	unique_tweets := map[string]bool{}

	for _, tweet := range tweet_data.Tweets {

		formatted_tweet_text := strings.Split(tweet.Text, "https://t.co")[0]

		if !unique_tweets[formatted_tweet_text] {
			unique_tweets[formatted_tweet_text] = true

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
}
