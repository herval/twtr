package util

import "errors"

var (
	ErrNoConsumerKey = errors.New("No consumer key and secret provided")
	ErrNoAccessToken = errors.New("No access token configured")
)
