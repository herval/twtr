package main

import "github.com/nsf/termbox-go"

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

func (t *TweetListController) OnKeyPress(key termbox.Event) EventResult {
	switch key.Ch {
	case 'e':
		t.View.Clear()
		t.Client.LoadTimeline()
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
