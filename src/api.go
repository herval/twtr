package main

import (
	"flag"
	"fmt"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"time"
)

type Client struct {
	TimelineTweets chan twitter.Tweet
	client         *twitter.Client
}

func NewClient() Client {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// TODO signin

	// Twitter client
	c := Client{
		TimelineTweets: make(chan twitter.Tweet, 100),
		client:         twitter.NewClient(httpClient),
	}
	return c
}

func (c *Client) Start() {
	go func() {
		for {
			tweets, err := c.GetTimeline()
			if err != nil {
				// TODO fail softer?
				Log.Println("Api error! ", err)
			} else {
				for _, t := range tweets {
					Log.Printf("Got Tweet:", t)
					c.TimelineTweets <- t
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

func (c *Client) GetTimeline() ([]twitter.Tweet, error) {
	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	tweets, _, err := c.client.Timelines.HomeTimeline(homeTimelineParams)

	if err != nil {
		return nil, err
	}

	return tweets, nil
}
