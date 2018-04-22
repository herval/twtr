package main

import termbox "github.com/nsf/termbox-go"

type Area struct {
	x0 int
	x1 int
	y0 int
	y1 int
}

func drawRepeat(startX int, endX int, startY int, endY int, char rune) {
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			draw(i, j, char)
		}
	}
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

func renderText(x int, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorBlack, termbox.ColorDefault)
		x++
	}
}

func renderTextHighlighted(x int, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorBlack|termbox.AttrReverse, termbox.ColorDefault)
		x++
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
