package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"strings"
	"fmt"
)

type Section interface {
	Draw(area Area)
	MinHeight(containerDimensions *Dimensions) int
}

// ---------------------------

type DefaultHeader struct {
	Text string
}

func (h *DefaultHeader) Draw(area Area) {
	renderText(area.x0, area.y0, "TWTR", false)
	renderTextRightJustified(area.x1-1, area.y0, h.Text)
}

func (h *DefaultHeader) MinHeight(containerDimensions *Dimensions) int {
	return 1
}

// ---------------------------

type Tweet struct {
	Window       *Window
	Content      *twitter.Tweet
	Highlighted  bool
	FullyVisible bool
}

func (t *Tweet) Draw(area Area) {
	y := area.y0

	ts, _ := t.Content.CreatedAtTime()

	headerText := fmt.Sprintf("@%s · %s", t.Content.User.ScreenName, timeAgo(ts))

	additionals := []string{}
	if t.Content.RetweetCount > 0 {
		additionals = append(
			additionals,
			fmt.Sprintf("🔁%d", t.Content.RetweetCount),
		)
	}
	if t.Content.FavoriteCount > 0 {
		additionals = append(
			additionals,
			fmt.Sprintf("⭐%d️", t.Content.FavoriteCount),
		)
	}
	if len(additionals) > 0 {
		for _, txt := range additionals {
			headerText += " · " + txt
		}
		headerText += " "
	}

	renderTextHighlighted(area.x0+1, area.x1, y, headerText, t.Highlighted)
	y += 1
	if y >= area.y1 {
		return
	}

	textLines := splitText(t.Content.FullText, area.x1-area.x0-1)
	for _, l := range textLines {
		if y >= area.y1 {
			return
		}
		renderTextHighlighted(area.x0+1, area.x1, y, strings.TrimLeft(l, " "), t.Highlighted)
		y += 1
	}

	drawRepeat(area.x0+1, area.x1-1, y, y, '-')
	if y >= area.y1 {
		return
	}

	t.FullyVisible = true // items that aren't fully visible won't get this set so we can scroll when selecting them
}

func (t *Tweet) MinHeight(containerDimensions *Dimensions) int {
	return len(splitText(t.Content.FullText, containerDimensions.width-1)) + 2
}

// ---------------------------

type TweetList struct {
	Window        *Window
	Tweets        []*Tweet
	SelectedIndex int
	ScrollOffset  int
}

func NewTweetList() TweetList {
	return TweetList{
		Tweets:        []*Tweet{},
		SelectedIndex: 0,
		ScrollOffset:  0,
	}
}

func (h *TweetList) Clear() {
	h.Tweets = []*Tweet{}
	h.SelectedIndex = 0
	h.ScrollOffset = 0
}

func (h *TweetList) AddTweet(t twitter.Tweet) bool {
	if !contains(h.Tweets, t) {
		h.Tweets = append(
			h.Tweets,
			&Tweet{
				Content:      &t,
				Highlighted:  false,
				FullyVisible: false,
			},
		)
		return true
	}

	return false
}

func contains(t []*Tweet, n twitter.Tweet) bool {
	for _, e := range t {
		if e.Content.IDStr == n.IDStr {
			return true
		}
	}
	return false
}

func (h *TweetList) SelectNext() {
	deselectAll(h.Tweets)

	h.SelectedIndex = min(len(h.Tweets)-1, h.SelectedIndex+1)

	if len(h.Tweets) > 0 {
		if !h.Tweets[h.SelectedIndex].FullyVisible {
			h.ScrollOffset = min(len(h.Tweets)-1, h.ScrollOffset+1) // scroll one down to try and fit
		}
	}
}

func (h *TweetList) SelectPrevious() {
	deselectAll(h.Tweets)

	h.SelectedIndex = max(0, h.SelectedIndex-1)

	if len(h.Tweets) > 0 {
		if !h.Tweets[h.SelectedIndex].FullyVisible {
			h.ScrollOffset = max(0, h.ScrollOffset-1) // scroll one up?
		}
	}
}

func deselectAll(tweets []*Tweet) {
	for _, t := range tweets {
		t.Highlighted = false
	}
}

func (h *TweetList) Draw(area Area) {
	hideAll(h.Tweets)
	availableArea := Area{
		x0: area.x0,
		x1: area.x1,
		y0: area.y0,
		y1: area.y1,
	}

	for i, t := range h.Tweets[h.ScrollOffset:] {
		if availableArea.y0 >= area.y1 { // no more space to render anything
			return
		}
		if (i + h.ScrollOffset) == h.SelectedIndex {
			t.Highlighted = true
		}

		availableSpace := &Dimensions{
			width:  availableArea.x1 - availableArea.x0 - 1,
			height: availableArea.y1 - availableArea.y0,
		}

		t.Draw(availableArea)

		yOffset := t.MinHeight(availableSpace)
		availableArea.y0 += yOffset // move the available space upper bound down by one tweet
	}

}
func hideAll(tweets []*Tweet) {
	for _, t := range tweets {
		t.FullyVisible = false
	}
}

func (t *TweetList) MinHeight(containerDimensions *Dimensions) int {
	return -1
}

// ---------------------------

type TextArea struct {
	Window *Window
	Text   string
}

func (h *TextArea) Draw(area Area) {
	renderText(area.x0, area.y0, h.Text, false)
}

func (h *TextArea) MinHeight(containerDimensions *Dimensions) int {
	return 1
}

// ---------------------------
