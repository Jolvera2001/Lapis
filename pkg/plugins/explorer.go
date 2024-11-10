package plugins

import (
	"lapis-project/pkg/kernel"
	l "lapis-project/pkg/layout"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type FileExplorerPlugin struct {
	theme         *material.Theme
	sidebarPlugin *SidebarPlugin
	fileList      []string
	listItems     map[string]*widget.Clickable
}

func NewFileExplorerPlugin(theme *material.Theme) *FileExplorerPlugin {
	return &FileExplorerPlugin{
		theme:     theme,
		listItems: make(map[string]*widget.Clickable),
		fileList:  []string{"file1.txt", "file2.txt", "folder1"},
	}
}

func (p *FileExplorerPlugin) ID() string {
	return "core.fileExplorer"
}

func (p *FileExplorerPlugin) Dependencies() []string {
	return []string{
		"core.sidebar",
	}
}

func (p *FileExplorerPlugin) Initialize(api kernel.API) error {
	// getting sidebar plugin
	if plugin, ok := api.GetPlugin("core.sidebar").(*SidebarPlugin); ok {
		p.sidebarPlugin = plugin

		explorerIcon, err := widget.NewIcon(icons.FileFolder)
		if err != nil {
			return err
		}

		p.sidebarPlugin.AddView(l.SidebarView{
			ID:      "explorer",
			Icon:    explorerIcon,
			Title:   "Files",
			Content: p.explorerLayout,
		})
	}

	return nil
}

func (p *FileExplorerPlugin) Start() error {
	return nil
}

func (p *FileExplorerPlugin) Stop() error {
	return nil
}

func (p *FileExplorerPlugin) explorerLayout(gtx layout.Context) layout.Dimensions {
	// List layout
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		// Could add a search bar at the top
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Add some padding around the list
			return layout.UniformInset(unit.Dp(8)).Layout(gtx,
				material.Body2(p.theme, "Files").Layout,
			)
		}),
		// The file list
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			var children []layout.FlexChild

			// Create a button for each file
			for _, file := range p.fileList {
				// Ensure we have a clickable button for this file
				if p.listItems[file] == nil {
					p.listItems[file] = new(widget.Clickable)
				}

				btn := p.listItems[file]
				file := file // Capture for closure

				// Add the file item to our list
				children = append(children,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						// Create a clickable row for the file
						return material.Clickable(gtx, btn, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(8)).Layout(gtx,
								material.Body1(p.theme, file).Layout,
							)
						})
					}),
				)
			}

			// Layout all file items vertically
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
		}),
	)
}
