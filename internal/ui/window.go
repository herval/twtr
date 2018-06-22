package ui

import (
	termbox "github.com/nsf/termbox-go"
	"time"
)


type Window struct {
	Header         Section
	Body           Section
	Footer         Section
	Controller     Controller
	PendingRefresh []Section
}

func (w *Window) ScheduleRefresh(s Section) {
	w.PendingRefresh = uniq(append(w.PendingRefresh, s))
}

func startRefresher(w *Window) {
	go func() {
		for {
			if len(w.PendingRefresh) > 0 {
				w.Draw()
				w.PendingRefresh = []Section{}
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
	w.Body = s
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
	f.PendingRefresh = []Section{}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// TODO draw header, body and footer
	var width, height = termbox.Size()

	renderFrame(width, height)

	nextY := 0

	frame := &Dimensions{
		width: width,
		height: height,
	}

	if f.Header != nil {
		var headerEnd = f.Header.MinHeight(frame) + nextY + 1
		f.Header.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: nextY + 1,
			y1: headerEnd,
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
			x0: 1,
			x1: width - 2,
			y0: nextY,
			y1: bodyEnd,
		})
		nextY = bodyEnd + 1
		renderDivider(width, height, bodyEnd)
	}

	if f.Footer != nil {
		footerTop := height - f.Footer.MinHeight(frame) - 1

		f.Footer.Draw(Area{
			x0: 2,
			x1: width - 2,
			y0: footerTop,
			y1: height - 1,
		})
	}

	termbox.Flush()
}
