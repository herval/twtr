package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	InitLogger()

	config, err := LoadConfig()
	if err != nil {
		Log.Fatal(err)
	}

	api, err := NewApiClient(config)
	if err != nil {
		Log.Fatal(err)
	}

	client, err := NewCachedClient(api)
	if err != nil {
		// TODO connect new account
		Log.Fatal(err)
	}

	// user profile... for display and stuff
	user, err := client.GetUser()
	if err != nil {
		Log.Fatal("Could not connect to Twitter API: %s\n", err.Error())
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
	defer client.Finish()

	controller := TweetListController{
		Window: &window,
		View:   &tweetList,
		Client: client,
	}
	controller.Show()

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
			Log.Fatal(ev.Err)
		}
	}

	fmt.Printf("Bye! 👋🐦")
}
