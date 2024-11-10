package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	"lapis-project/pkg/kernel"
	l "lapis-project/pkg/layout"
)

func main() {
	k := kernel.NewKernel()

	if err := k.Initialize(); err != nil {
		fmt.Println("error initializing kernel")
		os.Exit(0)
	}
	//  making sure when the window is taken off, the kernel shuts down
	defer k.Shutdown()

	w := new(app.Window)
	go func() {
		th := material.NewTheme()

		sidebar := l.NewSideBar(th)

		// layoutManager := kernel.NewLayoutManager()

		var ops op.Ops

		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				if e.Err != nil {
					log.Fatal(e.Err)
				}
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				// layout application
				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return sidebar.Layout(gtx)
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						// Main content area
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							label := material.Body1(th, "Main Content Area")
							return label.Layout(gtx)
						})
					}),
				)

				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
