package plugins

import(
	"gioui.org/widget/material"
	l "lapis-project/pkg/layout"
) 

type FileExplorerPlugin struct {
	theme         *material.Theme
	sidebarPlugin *l
}
