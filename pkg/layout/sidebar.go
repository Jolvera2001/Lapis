package layout

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SideBarState struct {
	activeView       string
	isExpanded       bool
	expandedWidth    float32
	activityBarWidth float32
	viewButtons      map[string]*widget.Clickable
	expandAnim       float32
}

func NewSideBarState() *SideBarState {
	return &SideBarState{
		isExpanded:       true,
		expandedWidth:    300,
		activityBarWidth: 48,
		viewButtons:      make(map[string]*widget.Clickable),
		expandAnim:       1.0, // Start expanded
	}
}

type SidebarView struct {
	ID       string
	Icon     *widget.Icon
	Title    string
	Content  layout.Widget
	Priority int
}

type Sidebar struct {
	state     *SideBarState
	views     []SidebarView
	theme     *material.Theme
	toggleBtn widget.Clickable
}

func NewSideBar(theme *material.Theme) *Sidebar {
	return &Sidebar{
		state: NewSideBarState(),
		theme: theme,
	}
}

func (s *Sidebar) AddView(view SidebarView) {
	if s.state.viewButtons[view.ID] == nil {
		s.state.viewButtons[view.ID] = new(widget.Clickable)
	}
	s.views = append(s.views, view)
}

func (s *Sidebar) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.layoutActivityBar(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if !s.state.isExpanded {
				return layout.Dimensions{}
			}
			return s.layoutSidebarPanel(gtx)
		}),
	)
}

func (s *Sidebar) layoutActivityBar(gtx layout.Context) layout.Dimensions {
	// Set the width constraint first
    width := int(s.state.activityBarWidth)
    gtx.Constraints.Min.X = width
    gtx.Constraints.Max.X = width

    return layout.UniformInset(unit.Dp(4)).Layout(gtx,
        func(gtx layout.Context) layout.Dimensions {
            gtx.Constraints.Min.X = int(s.state.activityBarWidth)
            gtx.Constraints.Max.X = int(s.state.activityBarWidth)

			// Paint the background for activity bar
            background := color.NRGBA{R: 0x1E, G: 0x1E, B: 0x2E, A: 0xFF} // Dark blue-ish background
            paint.FillShape(gtx.Ops,
                background,
                clip.Rect{
                    Max: image.Point{
                        X: width,
                        Y: gtx.Constraints.Max.Y,
                    },
                }.Op(),
            )

            // vertical list of icons
            return layout.Flex{
                Axis: layout.Vertical,
            }.Layout(gtx, // Notice this comma
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    var children []layout.FlexChild

                    for _, view := range s.views {
                        view := view
                        btn := s.state.viewButtons[view.ID]

                        children = append(children,
                            layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                                // check if this view is active
                                isActive := s.state.activeView == view.ID

								if isActive {
									background := color.NRGBA{R: 0x33, G: 0x33, B: 0xEE, A: 0xFF} // Blue color
									paint.FillShape(gtx.Ops,
										background,
										clip.Rect{
											Max: image.Point{
												X: gtx.Constraints.Max.X,
												Y: gtx.Dp(32),
											},
										}.Op(),
									)
								}

                                // style btn based on state
                                return material.Clickable(gtx, btn, func(gtx layout.Context) layout.Dimensions {
                                    // icon button
                                    size := gtx.Dp(32)
                                    gtx.Constraints.Min = image.Point{X: size, Y: size}
                                    gtx.Constraints.Max = image.Point{X: size, Y: size}

                                    if btn.Clicked(gtx) {
                                        if s.state.activeView == view.ID {
                                            // toggle if clicking active view
                                            s.state.isExpanded = !s.state.isExpanded
                                        } else {
                                            // switch to new view
                                            s.state.activeView = view.ID
                                            s.state.isExpanded = true
                                        }
                                    }

                                    // draw icon
                                    return view.Icon.Layout(gtx, s.theme.Fg)
                                })
                            }),
                            layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout), // Fixed Rigid spelling
                        )
                    }

                    return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
                }),
            )
        },
    )
}

func (s *Sidebar) layoutSidebarPanel(gtx layout.Context) layout.Dimensions {
	// set width for expanded sidebar
	gtx.Constraints.Min.X = int(s.state.expandedWidth)
	gtx.Constraints.Max.X = int(s.state.expandedWidth)

	return layout.UniformInset(unit.Dp(8)).Layout(gtx, 
		func(gtx layout.Context) layout.Dimensions {
			// find active view
			var activeView *SidebarView
			for _, view := range s.views {
				if view.ID == s.state.activeView {
					activeView = &view
					break
				}
			}

			if activeView == nil {
				return layout.Dimensions{}
			}

			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					title := material.H6(s.theme, activeView.Title)
					return layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, title.Layout)
				}),
				layout.Flexed(1, activeView.Content),
			)
		},
	)
}
