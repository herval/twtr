package main

import termbox "github.com/nsf/termbox-go"

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

type Window struct {
	header Section
	body   Section
	footer Section
}

func (f *Window) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (f *Window) Close() {
	termbox.Close()
}

func (f *Window) Draw() {
	// TODO draw header, body and footer
	var width, height = termbox.Size()

	renderFrame(width, height)

	var nextY = 0

	if f.header != nil {
		var headerEnd = f.header.Height(width, height) + nextY + 1
		f.header.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY + 1,
			y1: headerEnd,
		})
		nextY = headerEnd + 1
		renderDivider(width, height, headerEnd)
	}

	if f.body != nil {
		var bodyEnd = f.body.Height(width, height) + nextY + 1
		f.body.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY + 1,
			y1: bodyEnd,
		})
		nextY = nextY + 1
		renderDivider(width, height, bodyEnd)
	}

	if f.footer != nil {
		var footerEnd = f.footer.Height(width, height) + nextY + 1
		f.footer.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY + 1,
			y1: footerEnd,
		})
	}

	termbox.Flush()
}
