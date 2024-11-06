package kernel

type API interface {
	Subscribe(eventName string, handler EventHandler)
	Emit(event Event) error
	AddUIPlug(pluginId string, uiPlug UIPlug) error
}
