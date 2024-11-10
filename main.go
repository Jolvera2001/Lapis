package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"

	"lapis-project/pkg/kernel"
	l "lapis-project/pkg/layout"
)

func main() {
	// big bosses
	k := kernel.NewKernel()
	layoutManager := kernel.NewLayoutManager()
	th := material.NewTheme()

	// main ui
	sidebar := l.NewSideBar(th)

	// adding sidebar
	err := layoutManager.AddWidget(kernel.UIPlug{
		UI:          sidebar.Layout,
		Destination: "sidebar",
	})
	if err != nil {
		log.Printf("Failed to add sidebar widget: %v", err)
		os.Exit(1)
	}

	// initializing kernel
	if err := k.Initialize(); err != nil {
		fmt.Println("error initializing kernel")
		os.Exit(0)
	}
	//  making sure when the window is taken off, the kernel shuts down
	defer k.Shutdown()

	// Create window
	window := new(app.Window)

	// Start UI event loop
	go func() {
		var ops op.Ops

		for {
			e := window.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				if e.Err != nil {
					log.Fatal(e.Err)
				}
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				// Use layout manager to handle all layout
				layoutManager.Layout(gtx)

				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
