package kernel

import "gioui.org/layout"

type API interface {
	Subscribe(eventName string, handler EventHandler)
	Emit(event Event) error
	AddUIPlug(widget layout.Widget) error
}
