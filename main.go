package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"lapis-project/internal/utils"
)

func main() {
	go func() {
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
	myApp := utils.NewApp()
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			myApp.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
