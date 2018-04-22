package main

import "github.com/nsf/termbox-go"

func main() {
	InitLogger()
	var client = NewClient()

	var tweetList = NewTweetList()

	var window = Window{
		header:     &DefaultHeader{Text: "h"},
		body:       &tweetList,
		footer:     &DefaultHeader{Text: "f"},
		controller: nil,
	}
	window.Init()
	defer window.Close()

	controller := TweetListController{
		Window: &window,
		View:   &tweetList,
		Client: &client,
	}
	controller.Show()
	client.Start()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				return
			}

			if window.controller != nil && window.controller.OnKeyPress(ev) {
				window.Draw()
			}

		case termbox.EventResize:
			window.Draw()

		case termbox.EventMouse:
			// TODO mouse support?

		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
