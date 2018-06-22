package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"time"
)

type ApiClient struct {
	client         *twitter.Client
}

func NewApiClient(config *Config) (Client, error) {
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
	c := ApiClient{
		client:         twitter.NewClient(httpClient),
	}
	return &c, nil
}

func (c *ApiClient) GetUser() (*twitter.User, error) {
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

func (c *ApiClient) GetTimeline() (*TweetSet, error) {
	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     30,
		TweetMode: "extended",
	}
	tweets, _, err := c.client.Timelines.HomeTimeline(homeTimelineParams)

	if err != nil {
		return nil, err
	}

	return &TweetSet{
		Tweets:    tweets,
		UpdatedAt: time.Now(),
	}, nil
}
