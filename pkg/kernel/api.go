package kernel

type API interface {
	Subscribe(eventName string, handler EventHandler)
	Emit(event Event) error
}

