package tweet_view

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/herval/twtr/internal/ui"
)

type TweetView struct {
	Tweet *twitter.Tweet
}

func NewTweetView(t *twitter.Tweet) *TweetView {
	return &TweetView{
		Tweet: t,
	}
}

func (h *TweetView) Draw(area ui.Area) {
	availableArea := ui.Area{
		X0: area.X0,
		X1: area.X1,
		Y0: area.Y0,
		Y1: area.Y1,
	}

	body := ui.TextArea{
		Text: h.Tweet.FullText,
	}
	body.Draw(availableArea)
}

func (t *TweetView) MinHeight(containerDimensions *ui.Dimensions) int {
	return -1
}
