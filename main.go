package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Narven/incense/ui"
)

func main() {
	go func() {
		app.Size(unit.Dp(400), unit.Dp(600))
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			u := ui.UI{
				Gtx:   gtx,
				Theme: theme,
			}

			u.Text("Hello Incense")

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
