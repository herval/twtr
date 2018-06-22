package api

import (
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

type Client interface {
	GetUser() (*twitter.User, error)
	GetTimeline() (*TweetSet, error)
}

type TweetSet struct {
	UpdatedAt time.Time
	Tweets    []twitter.Tweet
}
