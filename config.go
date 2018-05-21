package main

import (
	"flag"
	"github.com/coreos/pkg/flagutil"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"errors"
)

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func LoadConfig() (*Config, error) {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	envFile := fmt.Sprint("%s/%s", dir, "env")

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Print("No env file found")
		flagutil.SetFlagsFromEnv(flags, "TWITTER")
	} else {
		flagutil.SetFlagsFromEnvFile(flags, "TWITTER", envFile)
	}

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		return nil, errors.New("Consumer key/secret and Access token/secret required")
	}

	return &Config{
		ConsumerKey:    *consumerKey,
		ConsumerSecret: *consumerSecret,
		AccessToken:    *accessToken,
		AccessSecret:   *accessSecret,
	}, nil
}
