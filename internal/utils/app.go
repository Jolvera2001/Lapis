package utils

import (
	"image/color"
	"lapis-project/internal/page"
	"lapis-project/internal/state"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

type App struct {
	state *state.AppState
	main  *page.MainPage
}

func NewApp() *App {
	state := &state.AppState{
		CurrentPage: state.MainPageID,
		DarkMode:    true,
		Theme:       material.NewTheme(),
	}
	state.UpdateTheme()

	return &App{
		state: state,
		main:  page.NewMainPage(state),
	}
}

func (a *App) Layout(gtx layout.Context) layout.Dimensions {
	paint.Fill(gtx.Ops, a.getBackgroundColor())

	switch a.state.CurrentPage {
	case state.MainPageID:
		return a.main.Layout(gtx)
	default: 
		return a.main.Layout(gtx)
	}
}

func (a *App) getBackgroundColor() color.NRGBA {
    if a.state.DarkMode {
        return color.NRGBA{R: 0x12, G: 0x12, B: 0x12, A: 0xFF}
    }
    return color.NRGBA{R: 0xF8, G: 0xF9, B: 0xFA, A: 0xFF}
}
