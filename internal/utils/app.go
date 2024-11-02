package utils

import (
	"lapis-project/internal/page"
	"lapis-project/internal/state"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type App struct {
	state *state.AppState
	main  *page.MainPage
}

func NewApp() *App {
	state := &state.AppState{
		CurrentPage: state.MainPageID,
		DarkMode:    false,
		Theme:       material.NewTheme(),
	}

	return &App{
		state: state,
		main:  page.NewMainPage(state),
	}
}

func (a *App) Layout(gtx layout.Context) layout.Dimensions {
	switch a.state.CurrentPage {
	case state.MainPageID:
		return a.main.Layout(gtx)
	default: 
		return a.main.Layout(gtx)
	}
}
