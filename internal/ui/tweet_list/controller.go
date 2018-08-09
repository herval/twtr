package tweet_list

import (
	"time"

	"github.com/herval/twtr/internal/api"
	"github.com/herval/twtr/internal/ui"
	"github.com/herval/twtr/internal/ui/tweet_view"
	"github.com/herval/twtr/internal/util"
	termbox "github.com/nsf/termbox-go"
)

type TweetListController struct {
	Window *ui.Window
	View   *TweetList
	Client api.Client
}

func NewTweetListController(window *ui.Window, client api.Client) *TweetListController {
	view := NewTweetList()
	return &TweetListController{
		Window: window,
		View:   &view,
		Client: client,
	}
}

func (t *TweetListController) OnShow() {
	go func() {
		for {
			refreshTweets(t)
			time.Sleep(60 * time.Second)
		}
	}()
}

func (t *TweetListController) Body() ui.View {
	return t.View
}

func refreshTweets(t *TweetListController) {
	tweets, err := t.Client.GetTimeline()
	if err != nil {
		util.Log.Println(err)

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

func (t *TweetListController) OnKeyPress(key termbox.Event) ui.EventResult {
	switch key.Ch {
	case 'e':
		t.View.Clear()
		refreshTweets(t)
		return ui.RedrawRequired

	case 'Q':
		return ui.ExitRequested
	}

	switch key.Key {
	case termbox.KeyEnter:
		tweet := t.View.SelectedTweet()
		t.Window.Push(
			tweet_view.NewTweetViewController(tweet.Content, t.Window, t.Client, t),
		)
		return ui.RedrawRequired

	case termbox.KeyArrowDown:
		t.View.SelectNext()
		return ui.RedrawRequired

	case termbox.KeyArrowUp:
		t.View.SelectPrevious()
		return ui.RedrawRequired

	case termbox.KeyCtrlQ:
		return ui.ExitRequested

	default:
		return ui.Noop
	}
}
