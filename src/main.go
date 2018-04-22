package main

import "github.com/nsf/termbox-go"

func main() {
	var window = Window{
		header: &DefaultHeader{},
		body:   nil,
		footer: nil,
	}
	window.Init()
	defer window.Close()

	window.Draw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				return
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			// TODO

			termbox.Flush()
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			window.Draw()

			termbox.Flush()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			// TODO

			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
