package ui

import (
	"fmt"
	"strings"

	"github.com/herval/twtr/internal/util"

	"github.com/dghubble/go-twitter/twitter"
)

type View interface {
	Draw(area Area)
	MinHeight(containerDimensions *Dimensions) int
}

// ---------------------------

type DefaultHeader struct {
	Text string
}

func (h *DefaultHeader) Draw(area Area) {
	renderText(area.X0, area.Y0, "TWTR", false)
	renderTextRightJustified(area.X1-1, area.Y0, h.Text)
}

func (h *DefaultHeader) MinHeight(containerDimensions *Dimensions) int {
	return 1
}

// ---------------------------

type Tweet struct {
	Content      *twitter.Tweet
	Highlighted  bool
	FullyVisible bool
}

func (t *Tweet) Draw(area Area) {
	y := area.Y0

	ts, _ := t.Content.CreatedAtTime()

	headerText := fmt.Sprintf("@%s · %s", t.Content.User.ScreenName, util.TimeAgo(ts))

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

	renderTextHighlighted(area.X0+1, area.X1, y, headerText, t.Highlighted)
	y += 1
	if y >= area.Y1 {
		return
	}

	textLines := util.SplitText(t.Content.FullText, area.X1-area.X0-1)
	for _, l := range textLines {
		if y >= area.Y1 {
			return
		}
		renderTextHighlighted(area.X0+1, area.X1, y, strings.TrimLeft(l, " "), t.Highlighted)
		y += 1
	}

	drawRepeat(area.X0+1, area.X1-1, y, y, '-')
	if y >= area.Y1 {
		return
	}

	t.FullyVisible = true // items that aren't fully visible won't get this set so we can scroll when selecting them
}

func (t *Tweet) MinHeight(containerDimensions *Dimensions) int {
	return len(util.SplitText(t.Content.FullText, containerDimensions.Width-1)) + 2
}

// ---------------------------

type TextArea struct {
	Text string
}

func (h *TextArea) Draw(area Area) {
	renderText(area.X0, area.Y0, h.Text, false)
}

func (h *TextArea) MinHeight(containerDimensions *Dimensions) int {
	return 1
}

// ---------------------------
