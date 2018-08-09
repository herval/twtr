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
}

func NewTweetViewController(tweet *twitter.Tweet, window *ui.Window, api api.Client) *TweetViewController {
	view := NewTweetView(tweet)
	// TODO fetch tweet details?

	return &TweetViewController{
		Tweet:  tweet,
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

func (t *TweetViewController) OnKeyPress(key termbox.Event) ui.EventResult {
	switch key.Ch {

	default:
		return ui.Noop
	}
}
