package kernel

import (
	"errors"
	"fmt"
	"sync"
)

type Kernel struct {
	uiPlugs  map[string][]UIPlug
	handlers map[string][]EventHandler
	plugins  map[string]Plugin
	mu       sync.Mutex
	started  bool
}

func NewKernel() *Kernel {
	return &Kernel{
		uiPlugs: make(map[string][]UIPlug),
		handlers: make(map[string][]EventHandler),
		plugins:  make(map[string]Plugin),
	}
}

func (k *Kernel) Emit(event Event) error {
	k.mu.Lock()
	handlers := k.handlers[event.Name]
	k.mu.Unlock()

	for _, handler := range handlers {
		if err := handler(event); err != nil {
			return err
		}
	}
	return nil
}

func (k *Kernel) Subscribe(eventName string, handler EventHandler) {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.handlers[eventName] == nil {
		k.handlers[eventName] = make([]EventHandler, 0)
	}

	k.handlers[eventName] = append(k.handlers[eventName], handler)
}

func (k *Kernel) AddUIPlug(pluginId string, uiPlug UIPlug) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.uiPlugs[pluginId] == nil {
		k.uiPlugs[pluginId] = make([]UIPlug, 0)
	}

	k.uiPlugs[pluginId] = append(k.uiPlugs[pluginId], uiPlug)
	return nil
}

func (k *Kernel) GetPlugin(id string) Plugin {
	k.mu.Lock()
	defer k.mu.Unlock()

	return k.plugins[id]
}

func (k *Kernel) Register(p Plugin) error {
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

func (k *Kernel) Initialize() error {
	k.mu.Lock()
	if k.started {
		return errors.New("kernel already started")
	}

	// sort plugins
	sorted, err := k.sortPluginByDependencies()
	if err != nil {
		return fmt.Errorf("error with sorting dependencies: %s", err)
	}

	k.mu.Unlock()

	// init plugins
	for id, p := range sorted {
		if err := p.Initialize(k); err != nil {
			return fmt.Errorf("failed to initialize plugin %d: %w", id, err)
		}
	}

	k.mu.Lock()
	// start plugins
	for id, p := range k.plugins {
		if err := p.Start(); err != nil {
			return fmt.Errorf("failed to start plugin %s: %w", id, err)
		}
	}

	k.started = true
	k.mu.Unlock()
	return nil
}

func (k *Kernel) Shutdown() error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if !k.started {
		return errors.New("Kernel not started")
	}

	// stop plugins
	for _, p := range k.plugins {
		if err := p.Stop(); err != nil {
			return err
		}
	}

	return nil
}

func (k *Kernel) sortPluginByDependencies() ([]Plugin, error) {
	// building graph
	graph := make(map[string][]string)
	pluginMap := make(map[string]Plugin)

	// building graph map
	for _, p := range k.plugins {
		graph[p.ID()] = p.Dependencies()
		pluginMap[p.ID()] = p
	}

	// track visited and temp for cycles
	visited := make(map[string]bool)
	temp := make(map[string]bool)

	// result
	var sorted []string

	// dfs
	var visit func(id string) error
	visit = func(id string) error {
		// check for cycle
		if temp[id] {
			return fmt.Errorf("cycle detected in plugin dependencies involving %s", id)
		}

		// skip if already visited
		if visited[id] {
			return nil
		}

		// mark temp
		temp[id] = true

		// visit dependencies
		for _, dep := range graph[id] {
			if _, exists := graph[dep]; !exists {
				return fmt.Errorf("plugin %s depends on non-existent plugin %s", id, dep)
			}

			if err := visit(dep); err != nil {
				return err
			}
		}

		// remove temp mark
		temp[id] = false
		//mark as visited
		visited[id] = true
		// add to sorted list
		sorted = append(sorted, id)

		return nil
	}

	// visit all plugins
	for id := range graph {
		if !visited[id] {
			if err := visit(id); err != nil {
				return nil, err
			}
		}
	}

	// convert sorted names back to array
	result := make([]Plugin, len(sorted))
	for i, id := range sorted {
		result[i] = pluginMap[id]
	}

	return result, nil
}
