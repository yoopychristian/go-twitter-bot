package main

import (
	"fmt"
	"go-twitter-bot/config"
	"go-twitter-bot/model"
	"log"
)

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	creds := model.Credentials{
		AccessToken:       config.ReadEnv("ACCESS_TOKEN", "your access token"),
		AccessTokenSecret: config.ReadEnv("ACCESS_TOKEN_SECRET", "your access token secret"),
		ConsumerKey:       config.ReadEnv("CONSUMER_KEY", "your consumer key"),
		ConsumerSecret:    config.ReadEnv("CONSUMER_SECRET", "your consumer secret"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := config.GetClient(creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	// Print out the pointer to our client
	// for now so it doesn't throw errors
	tweet, resp, err := client.Statuses.Update("go go @LFC", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

}
