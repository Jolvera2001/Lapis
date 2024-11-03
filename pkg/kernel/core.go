package kernel

import (
	"errors"
	"fmt"
	i "lapis-project/pkg/api/interfaces"
	"sync"
)

type Kernel struct {
	plugins map[string]i.Plugin
	mu      sync.Mutex
	started bool
}

func NewKernel() *Kernel {
	return &Kernel{
		plugins: make(map[string]i.Plugin),
	}
}

func (k *Kernel) Register(p i.Plugin) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.started {
		return errors.New("cannot register plugin: kernel already started")
	}

	id := p.ID()
	if _, exists := k.plugins[id]; exists {
		return fmt.Errorf("plugin %s already registered", id)
	}

	k.plugins[id] = p
	return nil
}

func (k *Kernel) Start() error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.started {
		return errors.New("kernel already started")
	}

	// init plugins
	for id, p := range k.plugins {
		if err := p.Initialize(); err != nil {
			return fmt.Errorf("failed to initialize plugin %s: %w", id, err)
		}
	}

	// start plugins
	for id, p := range k.plugins {
		if err := p.Start(); err != nil {
			return fmt.Errorf("failed to start plugin %s: %w", id, err)
		}
	}

	k.started = true
	return nil
}
