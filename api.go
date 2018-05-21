package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"time"
)

type Client struct {
	TimelineTweets chan twitter.Tweet
	client         *twitter.Client
}

func NewClient(config *Config) (*Client, error) {
	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		return nil, ErrNoConsumerKey
	}

	if config.AccessToken == "" || config.AccessSecret == "" {
		return nil, ErrNoAccessToken
	}

	conf := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := conf.Client(oauth1.NoContext, token)

	// TODO signin

	// Twitter client
	c := Client{
		TimelineTweets: make(chan twitter.Tweet, 100),
		client:         twitter.NewClient(httpClient),
	}
	return &c, nil
}

func (c *Client) Start() {
	go func() {
		for {
			c.LoadTimeline()
			time.Sleep(60 * time.Second)
		}
	}()
}

// load the user timeline and post it on the Timelines chan
func (c *Client) LoadTimeline() {
	tweets, err := c.GetTimeline()
	if err != nil {
		// TODO fail softer?
		Log.Println("Api error! ", err)
	} else {
		for _, t := range tweets {
			//Log.Printf("Got Tweet:", t)
			c.TimelineTweets <- t
		}
	}
}

func (c *Client) GetUser() (*twitter.User, error) {
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := c.client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Client) GetTimeline() ([]twitter.Tweet, error) {
	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     20,
		TweetMode: "extended",
	}
	tweets, _, err := c.client.Timelines.HomeTimeline(homeTimelineParams)

	if err != nil {
		return nil, err
	}

	return tweets, nil
}
