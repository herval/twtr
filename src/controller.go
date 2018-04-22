package main

import termbox "github.com/nsf/termbox-go"

type Controller interface {
	// return true if the event was handled
	OnKeyPress(key termbox.Event) bool
}

type TweetListController struct {
	Window *Window
	View   *TweetList
	Client *Client
}

func (t *TweetListController) Show() {
	go func() {
		for {
			select {
			case tweet := <-t.Client.TimelineTweets:
				t.View.AddTweet(&tweet)
				t.Window.ScheduleRefresh(t.View)
			}
		}
	}()

	t.Window.controller = t
	t.Window.SetBody(t.View)
}

func (t *TweetListController) OnKeyPress(key termbox.Event) bool {
	if key.Key == termbox.KeyArrowDown {
		t.View.SelectNext()
		return true
	}
	if key.Key == termbox.KeyArrowUp {
		t.View.SelectPrevious()
		return true
	}

	return false
}
