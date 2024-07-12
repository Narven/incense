package ui

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type UI struct {
	Gtx   layout.Context
	Theme *material.Theme
}

func (u *UI) Text(value string) {
	// Define an large label with an appropriate text:
	title := material.H1(u.Theme, value)

	// Change the color of the label.
	maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	title.Color = maroon

	// Change the position of the label.
	title.Alignment = text.Middle

	// Draw the label to the graphics context.
	title.Layout(u.Gtx)
}
