package kernel

import (
	"context"
	"time"
)

type Event struct {
	ID        string
	Name      string
	Data      interface{}
	Timestamp time.Time
	Context   context.Context
	Metadata  map[string]interface{}
}

type EventHandler func(Event) error
