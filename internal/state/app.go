package state

import (
	"image/color"

	"gioui.org/widget/material"
)

type Page int

const (
	MainPageID Page = iota
)

type AppState struct {
	CurrentPage Page
	DarkMode    bool
	Theme       *material.Theme
}

func (s *AppState) UpdateTheme() {
	th := material.NewTheme()
	if s.DarkMode {
		th.Palette = material.Palette{
			Bg:         color.NRGBA{R: 0x12, G: 0x12, B: 0x12, A: 0xFF},   // Dark background
            Fg:         color.NRGBA{R: 0xE0, G: 0xE0, B: 0xE0, A: 0xFF},   // Light text
            ContrastBg: color.NRGBA{R: 0x1F, G: 0x1F, B: 0x1F, A: 0xFF},   // Slightly lighter dark
            ContrastFg: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},   // White text
		}
	}
	s.Theme = th
}