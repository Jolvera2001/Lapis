package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"lapis-project/pkg/kernel"
	"lapis-project/pkg/plugins"
	"lapis-project/pkg/theme"
)

func main() {
	// big bosses
	k := kernel.NewKernel()
	layoutManager := kernel.NewLayoutManager()
	th, nil := theme.NewCustomTheme()

	// main ui
	sidebarPlugin := plugins.NewSidebarPlugin(th.Theme)
	if err := k.Register(sidebarPlugin); err != nil {
		log.Printf("Failed to register sidebar plugin: %v", err)
		os.Exit(1)
	}

	// real plugins
	explorerPlugin := plugins.NewFileExplorerPlugin(th.Theme)
	if err := k.Register(explorerPlugin); err != nil {
		log.Printf("Failed to register plugin: %v", err)
		os.Exit(1)
	}

	if err := k.AddUIPlug("core.layout", kernel.UIPlug{
		UI: layoutManager.Layout,
		Destination: "root",
	}); err != nil {
		log.Printf("Failed to add layout manager: %v", err)
		os.Exit(1)
	}

	// initializing kernel
	if err := k.Initialize(); err != nil {
		fmt.Println("error initializing kernel")
		os.Exit(0)
	}
	//  making sure when the window is taken off, the kernel shuts down
	defer k.Shutdown()

	k.ConnectLayoutManager(layoutManager)

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
