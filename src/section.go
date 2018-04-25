package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"strings"
	"fmt"
)

type Section interface {
	Draw(area Area)
	MinHeight(windowWidth int, windowHeigh int) int
}

// ---------------------------

type DefaultHeader struct {
	Text string
}

func (h *DefaultHeader) Draw(area Area) {
	renderText(area.x0, area.y0, "TWTR")
	renderTextRightJustified(area.x1-1, area.y0, h.Text)
}

func (h *DefaultHeader) MinHeight(windowWidth int, windowHeigh int) int {
	return 1
}

// ---------------------------

type TweetList struct {
	Window        *Window
	Tweets        []*twitter.Tweet
	SelectedIndex int
}

func NewTweetList() TweetList {
	return TweetList{
		Tweets:        []*twitter.Tweet{},
		SelectedIndex: 0,
	}
}

func (h *TweetList) AddTweet(t *twitter.Tweet) {
	if !contains(h.Tweets, t) {
		h.Tweets = append(h.Tweets, t)
	}
}

func contains(t []*twitter.Tweet, n *twitter.Tweet) bool {
	for _, e := range t {
		if e.IDStr == n.IDStr {
			return true
		}
	}
	return false
}

func (h *TweetList) SelectNext() {
	h.SelectedIndex += 1
	if h.SelectedIndex >= len(h.Tweets) {
		h.SelectedIndex = 0
	}
}

func (h *TweetList) SelectPrevious() {
	h.SelectedIndex -= 1
	if h.SelectedIndex < 0 {
		h.SelectedIndex = len(h.Tweets) - 1
	}
}

func (h *TweetList) Draw(area Area) {
	y := area.y0
	for i, t := range h.Tweets {
		if y >= area.y1 {
			return
		}
		highlighted := i == h.SelectedIndex

		ts, _ := t.CreatedAtTime()

		renderTextHighlighted(area.x0, y, fmt.Sprintf("@%s · %s", t.User.ScreenName, timeAgo(ts)), highlighted)
		y += 1
		if y >= area.y1 {
			return
		}

		textLines := splitText(t.FullText, area.x1-area.x0)
		for _, l := range textLines {
			renderTextHighlighted(area.x0, y, strings.TrimLeft(l, " "), highlighted)
			y += 1
			if y >= area.y1 {
				return
			}
		}

		drawRepeat(area.x0, area.x1, y, y, '-')
		y += 1
	}

}

func (t *TweetList) MinHeight(windowWidth int, windowHeigh int) int {
	return -1
}

// ---------------------------

type TextArea struct {
	Window *Window
	Text   string
}

func (h *TextArea) Draw(area Area) {
	renderText(area.x0, area.y0, h.Text)
}

func (h *TextArea) MinHeight(windowWidth int, windowHeigh int) int {
	return 1
}

// ---------------------------
