package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	InitLogger()

	config := LoadConfig()
	client, err := NewClient(config)
	if err != nil {
		// TODO connect new account
		panic(err)
	}

	// user profile... for display and stuff
	user, err := client.GetUser()
	if err != nil {
		fmt.Printf("Could not connect to Twitter API: %s\n", err.Error())
		os.Exit(1)
	}

	tweetList := NewTweetList()

	window := Window{
		header:     &DefaultHeader{Text: user.Name},
		body:       &tweetList,
		footer:     &TextArea{Text: "(N)ew Tweet | (R)eply | (F)avorite | (Q)uit"},
		controller: nil,
	}
	window.Init()
	defer window.Close()

	controller := TweetListController{
		Window: &window,
		View:   &tweetList,
		Client: client,
	}
	controller.Show()
	client.Start()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				return
			}

			if window.controller != nil && window.controller.OnKeyPress(ev) {
				window.Draw()
			}

		case termbox.EventResize:
			window.Draw()

		case termbox.EventMouse:
			// TODO mouse support?

		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
