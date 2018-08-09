package ui

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	Header         View
	Body           View
	Footer         View
	controller     Controller
	pendingRefresh []View
}

func (w *Window) ScheduleRefresh(s View) {
	w.pendingRefresh = uniq(append(w.pendingRefresh, s))
}

func startRefresher(w *Window) {
	go func() {
		for {
			if len(w.pendingRefresh) > 0 {
				w.Draw()
				w.pendingRefresh = []View{}
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func (w *Window) CurrentController() Controller {
	return w.controller
}

func (f *Window) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	startRefresher(f)

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (w *Window) Push(c Controller) {
	w.controller = c
	w.Body = c.Body()
	w.ScheduleRefresh(w.Body)
	c.OnShow()
}

func (f *Window) Close() {
	termbox.Close()
}

func uniq(Views []View) []View {
	keys := make(map[View]bool)
	list := []View{}
	for _, entry := range Views {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (f *Window) Draw() {
	f.pendingRefresh = []View{}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// TODO draw header, body and footer
	width, height := termbox.Size()

	renderFrame(width, height)

	nextY := 0

	frame := &Dimensions{
		Width:  width,
		Height: height,
	}

	if f.Header != nil {
		var headerEnd = f.Header.MinHeight(frame) + nextY + 1
		f.Header.Draw(Area{
			X0: 2,
			X1: width - 2,
			Y0: nextY + 1,
			Y1: headerEnd,
		})
		nextY = headerEnd + 1
		renderDivider(width, height, headerEnd)
	}

	// body takes all available space
	if f.Body != nil {
		var bodyEnd = height - 1
		if f.Footer != nil {
			bodyEnd = height - f.Footer.MinHeight(frame) - 2
		}

		f.Body.Draw(Area{
			X0: 1,
			X1: width - 2,
			Y0: nextY,
			Y1: bodyEnd,
		})
		nextY = bodyEnd + 1
		renderDivider(width, height, bodyEnd)
	}

	if f.Footer != nil {
		footerTop := height - f.Footer.MinHeight(frame) - 1

		f.Footer.Draw(Area{
			X0: 2,
			X1: width - 2,
			Y0: footerTop,
			Y1: height - 1,
		})
	}

	termbox.Flush()
}
