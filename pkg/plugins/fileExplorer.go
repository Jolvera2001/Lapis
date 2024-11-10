package plugins

import (
	"lapis-project/pkg/kernel"
	l "lapis-project/pkg/layout"

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
	return nil
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
