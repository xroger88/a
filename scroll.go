package main

import (
	mus "github.com/as/text/mouse"
	"github.com/as/ui/win"
	"golang.org/x/mobile/event/paint"
)

func doScrollEvent(act *win.Win, e mus.ScrollEvent) {
	if e.Button == -1 {
		e.Dy = -e.Dy
	}
	actTag.Body.Scroll(e.Dy)
	act.Window().Send(paint.Event{})
}
