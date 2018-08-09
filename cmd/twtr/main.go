package main

import (
	"fmt"

	"github.com/herval/twtr/internal/api"
	"github.com/herval/twtr/internal/ui"
	"github.com/herval/twtr/internal/ui/tweet_list"
	"github.com/herval/twtr/internal/util"
	"github.com/nsf/termbox-go"
)

func main() {
	util.InitLogger()

	config, err := util.LoadConfig()
	if err != nil {
		util.Log.Fatal(err)
	}

	apiClient, err := api.NewApiClient(config)
	if err != nil {
		util.Log.Fatal(err)
	}

	client, err := api.NewCachedClient(apiClient)
	if err != nil {
		// TODO connect new account
		util.Log.Fatal(err)
	}

	// user profile... for display and stuff
	user, err := client.GetUser()
	if err != nil {
		util.Log.Fatal("Could not connect to Twitter API: %s\n", err.Error())
	}

	window := ui.Window{
		Header: &ui.DefaultHeader{Text: user.Name},
		Footer: &ui.TextArea{Text: "(N)ew Tweet | (R)eply | (F)avorite | R(e)fresh | (Q)uit"},
	}
	window.Init()
	defer window.Close()
	defer client.Finish()

	window.Push(
		tweet_list.NewTweetListController(&window, client),
	)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if window.CurrentController() != nil {
				switch window.CurrentController().OnKeyPress(ev) {
				case ui.Noop:
					// nothing
				case ui.RedrawRequired:
					window.Draw()
				case ui.ExitRequested:
					return
				}
			}

		case termbox.EventResize:
			window.Draw()

		case termbox.EventMouse:
			// TODO mouse support?

		case termbox.EventError:
			util.Log.Fatal(ev.Err)
		}
	}

	fmt.Printf("Bye! 👋🐦")
}
