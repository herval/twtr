package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	InitLogger()

	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

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
		footer:     &TextArea{Text: "(N)ew Tweet | (R)eply | (F)avorite | R(e)fresh | (Q)uit"},
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
			if window.controller != nil {
				switch window.controller.OnKeyPress(ev) {
				case Noop:
					// nothing
				case RedrawRequired:
					window.Draw()
				case ExitRequested:
					return
				}
			}

		case termbox.EventResize:
			window.Draw()

		case termbox.EventMouse:
			// TODO mouse support?

		case termbox.EventError:
			panic(ev.Err)
		}
	}

	fmt.Printf("Bye! 👋🐦")
}
