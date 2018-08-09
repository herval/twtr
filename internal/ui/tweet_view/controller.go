package tweet_view

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/herval/twtr/internal/api"
	"github.com/herval/twtr/internal/ui"
	termbox "github.com/nsf/termbox-go"
)

type TweetViewController struct {
	Tweet  *twitter.Tweet
	Window *ui.Window
	View   ui.View
	Client api.Client
	Parent ui.Controller
}

func NewTweetViewController(
	tweet *twitter.Tweet,
	window *ui.Window,
	api api.Client,
	parent ui.Controller,
) *TweetViewController {
	view := NewTweetView(tweet)
	// TODO fetch tweet details?

	return &TweetViewController{
		Tweet:  tweet,
		Parent: parent,
		Window: window,
		Client: api,
		View:   view,
	}
}

func (t *TweetViewController) Body() ui.View {
	return t.View
}

func (t *TweetViewController) OnShow() {

}

func (t *TweetViewController) Close() {
	t.Window.Push(t.Parent)
}

func (t *TweetViewController) OnKeyPress(key termbox.Event) ui.EventResult {
	switch key.Ch {
	case 'Q':
		return ui.ExitRequested
	}

	switch key.Key {
	case termbox.KeyEsc:
		t.Close()
		return ui.RedrawRequired

	case termbox.KeyCtrlQ:
		return ui.ExitRequested

	default:
		return ui.Noop
	}
}
