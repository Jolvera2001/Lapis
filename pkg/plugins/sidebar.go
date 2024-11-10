package plugins

// THIS IS A MAIN UI COMPONENT
// THIS IS WRAPPED IN A PLUGIN TO EXPOSE INTERNAL
// UI SYSTEM DEFINED IN THE "layout" FOLDER

import (
	"lapis-project/pkg/kernel"
	l "lapis-project/pkg/layout"

	"gioui.org/widget/material"
)

type SidebarPlugin struct {
	sidebar *l.Sidebar
	theme   *material.Theme
}

func NewSidebarPlugin(theme *material.Theme) *SidebarPlugin {
	return &SidebarPlugin{
		theme:   theme,
		sidebar: l.NewSideBar(theme),
	}
}

func (p *SidebarPlugin) ID() string {
	return "core.sidebar"
}

func (p *SidebarPlugin) Dependencies() []string {
	return nil
}

func (p *SidebarPlugin) Initialize(api kernel.API) error {
	return api.AddUIPlug(p.ID(), kernel.UIPlug{
		UI:          p.sidebar.Layout,
		Destination: "sidebar",
	})
}

func (p *SidebarPlugin) Start() error {
	return nil
}

func (p *SidebarPlugin) Stop() error {
	return nil
}

func (p *SidebarPlugin) AddView(view l.SidebarView) {
	p.sidebar.AddView(view)
}
