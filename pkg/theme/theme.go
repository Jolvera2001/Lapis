package theme

import (
	_ "embed"
	"fmt"
	"image/color"

	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
)

//go:embed fonts/Nunito-VariableFont_wght.ttf
var nunitoVariable []byte

//go:embed fonts/Nunito-Italic-VariableFont_wght.ttf
var nunitoVariableItalic []byte

type CustomTheme struct {
	*material.Theme

	SidebarBg     color.NRGBA
	ActivityBarBg color.NRGBA
	ActiveItemBg  color.NRGBA
	IconColor     color.NRGBA
	ActiveIcon    color.NRGBA
}

func NewCustomTheme() (*CustomTheme, error) {
	// Parse regular variable font
	regularFace, err := opentype.Parse(nunitoVariable)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Nunito regular: %v", err)
	}

	// Parse italic variable font
	italicFace, err := opentype.Parse(nunitoVariableItalic)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Nunito italic: %v", err)
	}

	// Create font collection with different weights and styles
	fontCollection := []font.FontFace{
		// Regular weights
		{
			Font: font.Font{
				Typeface: "Nunito",
				Style:    font.Regular,
				Weight:   font.Normal,
			},
			Face: regularFace,
		},
		{
			Font: font.Font{
				Typeface: "Nunito",
				Style:    font.Regular,
				Weight:   font.Bold,
			},
			Face: regularFace,
		},
		// Italic weights
		{
			Font: font.Font{
				Typeface: "Nunito",
				Style:    font.Italic,
				Weight:   font.Normal,
			},
			Face: italicFace,
		},
		{
			Font: font.Font{
				Typeface: "Nunito",
				Style:    font.Italic,
				Weight:   font.Bold,
			},
			Face: italicFace,
		},
	}

	// Create base theme
	base := material.NewTheme()
	base.Shaper = text.NewShaper(text.WithCollection(fontCollection))

	return &CustomTheme{
		Theme: base,

		SidebarBg:     color.NRGBA{R: 0xE8, G: 0xE8, B: 0xE8, A: 0xFF},
		ActivityBarBg: color.NRGBA{R: 0xF0, G: 0xF0, B: 0xF0, A: 0xFF},
		ActiveItemBg:  color.NRGBA{R: 0x33, G: 0x33, B: 0xEE, A: 0xFF},
		IconColor:     color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		ActiveIcon:    color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	}, nil
}
