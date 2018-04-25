package main

import (
	termbox "github.com/nsf/termbox-go"
	"time"
)


type Window struct {
	header         Section
	body           Section
	footer         Section
	controller     Controller
	pendingRefresh []Section
}

func (w *Window) ScheduleRefresh(s Section) {
	w.pendingRefresh = uniq(append(w.pendingRefresh, s))
}

func startRefresher(w *Window) {
	go func() {
		for {
			if len(w.pendingRefresh) > 0 {
				w.Draw()
				w.pendingRefresh = []Section{}
			}
			time.Sleep(1 * time.Second)
		}
	}()
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

func (w *Window) SetBody(s Section) {
	w.body = s
	w.ScheduleRefresh(s)
}

func (f *Window) Close() {
	termbox.Close()
}

func uniq(sections []Section) []Section {
	keys := make(map[Section]bool)
	list := []Section{}
	for _, entry := range sections {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (f *Window) Draw() {
	f.pendingRefresh = []Section{}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// TODO draw header, body and footer
	var width, height = termbox.Size()

	renderFrame(width, height)

	var nextY = 0

	if f.header != nil {
		var headerEnd = f.header.MinHeight(width, height) + nextY + 1
		f.header.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY + 1,
			y1: headerEnd,
		})
		nextY = headerEnd + 1
		renderDivider(width, height, headerEnd)
	}

	// body takes all available space
	if f.body != nil {
		var bodyEnd = height - 1
		if f.footer != nil {
			bodyEnd = height - f.footer.MinHeight(width, height) - 2
		}

		f.body.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY,
			y1: bodyEnd,
		})
		nextY = bodyEnd + 1
		renderDivider(width, height, bodyEnd)
	}

	if f.footer != nil {
		footerTop := height - f.footer.MinHeight(width, height) - 1

		f.footer.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: footerTop,
			y1: height - 1,
		})
	}

	termbox.Flush()
}
