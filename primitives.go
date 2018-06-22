package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"fmt"
)

const (
	TopLeftBorder      = 0x2554
	TopRightBorder     = 0x2557
	BottomLeftBorder   = 0x255A
	BottomRightBorder  = 0x255D
	LeftDividerBorder  = 0x2560
	RightDividerBorder = 0x2563
	HorizontalBorder   = 0x2550
	VerticalLineBorder = 0x2551
)

type Area struct {
	x0 int
	x1 int
	y0 int
	y1 int
}

type Dimensions struct {
	width  int
	height int
}

func drawRepeat(startX int, endX int, startY int, endY int, char rune) {
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			draw(i, j, char)
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitText(text string, maxLineSize int) []string {
	res := []string{}

	i := 0
	for len(text) > 0 {
		l := text[0:min(maxLineSize, len(text))]
		res = append(res, l)
		i += maxLineSize
		text = text[min(maxLineSize, len(text)):]
	}

	return res
}

func timeAgo(ts time.Time) string {
	dur := time.Since(ts)
	if dur.Minutes() < 60 {
		return fmt.Sprintf("%dm", int(dur.Minutes()))
	}
	if dur.Hours() < 24 {
		return fmt.Sprintf("%dh", int(dur.Hours()))
	}
	return fmt.Sprintf("%dd", int(dur.Hours()/24))
}

func draw(x int, y int, char rune) {
	termbox.SetCell(x, y, char, termbox.ColorWhite, termbox.ColorDefault)
}

func renderFrame(width int, height int) {
	draw(0, 0, TopLeftBorder)
	draw(width-1, 0, TopRightBorder)
	draw(0, height-1, BottomLeftBorder)
	draw(width-1, height-1, BottomRightBorder)
	drawRepeat(1, width-2, 0, 0, HorizontalBorder)
	drawRepeat(1, width-2, height-1, height-1, HorizontalBorder)

	drawRepeat(0, 0, 1, height-2, VerticalLineBorder)
	drawRepeat(width-1, width-1, 1, height-2, VerticalLineBorder)
}

func renderDivider(width int, height int, dividerY int) {
	draw(0, dividerY, LeftDividerBorder)
	draw(width-1, dividerY, RightDividerBorder)
	drawRepeat(1, width-2, dividerY, dividerY, HorizontalBorder)
}

func renderTextRightJustified(xEnd int, yEnd int, msg string) {
	x := xEnd
	y := yEnd
	for _, c := range reverse(msg) {

		termbox.SetCell(x, y, c, termbox.ColorBlack, termbox.ColorDefault)
		x--
	}
}

// render and move the cursor
func renderText(x int, y int, msg string, highlighted bool) int {
	color := termbox.ColorBlack
	if highlighted {
		color |= termbox.AttrReverse
	}

	for _, c := range msg {
		termbox.SetCell(x, y, c, color, termbox.ColorDefault)
		x++
	}

	return x
}

func renderTextHighlighted(x0 int, xn int, y int, msg string, highlighted bool) {
	x := x0
	if highlighted {
		x = renderText(x, y, msg, true)
		for x < xn {
			termbox.SetCell(x, y, ' ', termbox.ColorBlack|termbox.AttrReverse, termbox.ColorDefault)
			x++
		}
	} else {
		x = renderText(x, y, msg, false)
		for x < xn {
			termbox.SetCell(x, y, ' ', termbox.ColorBlack, termbox.ColorDefault)
			x++
		}
	}
}

func reverse(s string) string {
	chars := []rune(s)
	rev := make([]rune, 0, len(chars))
	for i := len(chars) - 1; i >= 0; i-- {
		rev = append(rev, chars[i])
	}
	return string(rev)
}
