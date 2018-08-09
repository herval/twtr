package tweet_list

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/herval/twtr/internal/ui"
	"github.com/herval/twtr/internal/util"
)

type TweetList struct {
	Tweets        []*ui.Tweet
	SelectedIndex int
	ScrollOffset  int
}

func NewTweetList() TweetList {
	return TweetList{
		Tweets:        []*ui.Tweet{},
		SelectedIndex: 0,
		ScrollOffset:  0,
	}
}

func (h *TweetList) Clear() {
	h.Tweets = []*ui.Tweet{}
	h.SelectedIndex = 0
	h.ScrollOffset = 0
}

func (h *TweetList) AddTweet(t twitter.Tweet) bool {
	if !contains(h.Tweets, t) {
		h.Tweets = append(
			h.Tweets,
			&ui.Tweet{
				Content:      &t,
				Highlighted:  false,
				FullyVisible: false,
			},
		)
		return true
	}

	return false
}

func (h *TweetList) SelectedTweet() *ui.Tweet {
	return h.Tweets[h.SelectedIndex]
}

func contains(t []*ui.Tweet, n twitter.Tweet) bool {
	for _, e := range t {
		if e.Content.IDStr == n.IDStr {
			return true
		}
	}
	return false
}

func (h *TweetList) SelectNext() {
	deselectAll(h.Tweets)

	h.SelectedIndex = util.Min(len(h.Tweets)-1, h.SelectedIndex+1)

	if len(h.Tweets) > 0 {
		if !h.Tweets[h.SelectedIndex].FullyVisible {
			h.ScrollOffset = util.Min(len(h.Tweets)-1, h.ScrollOffset+1) // scroll one down to try and fit
		}
	}
}

func (h *TweetList) SelectPrevious() {
	deselectAll(h.Tweets)

	h.SelectedIndex = util.Max(0, h.SelectedIndex-1)

	if len(h.Tweets) > 0 {
		if !h.Tweets[h.SelectedIndex].FullyVisible {
			h.ScrollOffset = util.Max(0, h.ScrollOffset-1) // scroll one up?
		}
	}
}

func deselectAll(tweets []*ui.Tweet) {
	for _, t := range tweets {
		t.Highlighted = false
	}
}

func (h *TweetList) Draw(area ui.Area) {
	hideAll(h.Tweets)
	availableArea := ui.Area{
		X0: area.X0,
		X1: area.X1,
		Y0: area.Y0,
		Y1: area.Y1,
	}

	for i, t := range h.Tweets[h.ScrollOffset:] {
		if availableArea.Y0 >= area.Y1 { // no more space to render anything
			return
		}
		if (i + h.ScrollOffset) == h.SelectedIndex {
			t.Highlighted = true
		}

		availableSpace := &ui.Dimensions{
			Width:  availableArea.X1 - availableArea.X0 - 1,
			Height: availableArea.Y1 - availableArea.Y0,
		}

		t.Draw(availableArea)

		yOffset := t.MinHeight(availableSpace)
		availableArea.Y0 += yOffset // move the available space upper bound down by one tweet
	}

}
func hideAll(tweets []*ui.Tweet) {
	for _, t := range tweets {
		t.FullyVisible = false
	}
}

func (t *TweetList) MinHeight(containerDimensions *ui.Dimensions) int {
	return -1
}
