package kernel

import (
	"sync"

	"gioui.org/layout"
)

type LayoutManager struct {
	Widgets map[string][]UIPlug
	mu      sync.Mutex
}

func NewLayoutManager() *LayoutManager {
	return &LayoutManager{
		Widgets: make(map[string][]UIPlug),
	}
}

func (l *LayoutManager) AddWidget(plug UIPlug) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Widgets[plug.Destination] == nil {
		l.Widgets[plug.Destination] = make([]UIPlug, 0)
	}

	l.Widgets[plug.Destination] = append(l.Widgets[plug.Destination], plug)
	return nil
}

func (l *LayoutManager) Layout(gtx layout.Context) layout.Dimensions {
	l.mu.Lock()
	defer l.mu.Unlock()

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx, 
		// sidebar
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if widgets := l.Widgets["sidebar"]; len(widgets) > 0 {
				return widgets[0].UI(gtx)
			}
			return layout.Dimensions{}
		}),
		// any other areas
	)
}

