package page

import (
	"lapis-project/internal/state"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type MainPage struct {
	state *state.AppState
}

func NewMainPage(state *state.AppState) *MainPage {
	return &MainPage{
		state: state,
	}
}

func (p *MainPage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.H6(p.state.Theme, "Main Page").Layout(gtx)
		}),
	)
}
