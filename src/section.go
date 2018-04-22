package main

type Section interface {
	Draw(area Area)
	Height(windowWidth int, windowHeigh int) int
}

type DefaultHeader struct{}

func (h *DefaultHeader) Draw(area Area) {
	renderText(area.x0, area.y0, "TWTR")
	renderTextRightJustified(area.x1-1, area.y0, "@herval")
}

func (h *DefaultHeader) Height(windowWidth int, windowHeigh int) int {
	return 1
}
