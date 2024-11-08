package kernel

import (
	"fmt"
	"sync"

	"gioui.org/layout"
)

type ZoneConfig struct {
	Direction  layout.Direction
	MaxWidgets int
	Dimensions layout.Dimensions
	Priority   int
	Spacing    int
	Padding    int
}

type LayoutManager struct {
	zones             map[string][]UIPlug
	zoneConfigs       map[string]ZoneConfig
	validDestinations map[string]bool
	mu                sync.Mutex
}

func NewLayoutManager() *LayoutManager {
	return &LayoutManager{
		zones: make(map[string][]UIPlug),
		validDestinations: map[string]bool{
			"left-sidebar": true,
			"main-area":    true,
			"header":       true,
		},
	}
}

func (l *LayoutManager) AddWidget(plug UIPlug) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.validDestinations[plug.Destination] {
		return fmt.Errorf("invalid destination: %s", plug.Destination)
	}

	l.zones[plug.Destination] = append(l.zones[plug.Destination], plug)
	return nil
}

