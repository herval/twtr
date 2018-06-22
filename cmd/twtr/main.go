package main

import 	(
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/herval/twtr/internal/api"
	"github.com/herval/twtr/internal/util"
	"github.com/herval/twtr/internal/ui"

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

	tweetList := ui.NewTweetList()

	window := ui.Window{
		Header:     &ui.DefaultHeader{Text: user.Name},
		Body:       &tweetList,
		Footer:     &ui.TextArea{Text: "(N)ew Tweet | (R)eply | (F)avorite | R(e)fresh | (Q)uit"},
		Controller: nil,
	}
	window.Init()
	defer window.Close()
	defer client.Finish()

	controller := ui.TweetListController{
		Window: &window,
		View:   &tweetList,
		Client: client,
	}
	controller.Show()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if window.Controller != nil {
				switch window.Controller.OnKeyPress(ev) {
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
