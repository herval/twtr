package ui

import (
	"github.com/nsf/termbox-go"
)

type EventResult int

const (
	RedrawRequired EventResult = iota
	ExitRequested
	Noop
)

type Controller interface {
	// return true if the event was handled
	OnKeyPress(key termbox.Event) EventResult
	Body() View
	OnShow()
}
