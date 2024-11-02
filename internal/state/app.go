package state

import "gioui.org/widget/material"

type Page int

const (
	MainPageID Page = iota
)

type AppState struct {
	CurrentPage Page
	DarkMode    bool
	Theme       *material.Theme
}