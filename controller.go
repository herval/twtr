package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type EventResult int

const (
	RedrawRequired EventResult = iota
	ExitRequested
	Noop
)

type Controller interface {
	// return true if the event was handled
	OnKeyPress(key termbox.Event) EventResult
}

type TweetListController struct {
	Window *Window
	View   *TweetList
	Client Client
}

func (t *TweetListController) Show() {
	go func() {
		for {
			refreshTweets(t)
			time.Sleep(60 * time.Second)
		}
	}()

	t.Window.controller = t
	t.Window.SetBody(t.View)
}

func refreshTweets(t *TweetListController) {
	tweets, err := t.Client.GetTimeline()
	if err != nil {
		Log.Println(err)

	} else {
		refresh := false
		for _, tweet := range tweets.Tweets {
			refresh = t.View.AddTweet(tweet) || refresh
		}

		if refresh {
			//Log.Println("Refreshing...")
			t.Window.ScheduleRefresh(t.View)
		}
	}
}

func (t *TweetListController) OnKeyPress(key termbox.Event) EventResult {
	switch key.Ch {
	case 'e':
		t.View.Clear()
		refreshTweets(t)
		return RedrawRequired

	case 'Q':
		return ExitRequested
	}

	switch key.Key {
	case termbox.KeyArrowDown:
		t.View.SelectNext()
		return RedrawRequired

	case termbox.KeyArrowUp:
		t.View.SelectPrevious()
		return RedrawRequired

	case termbox.KeyCtrlQ:
		return ExitRequested

	default:
		return Noop
	}
}
