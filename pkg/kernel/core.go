package kernel

import (
	i "lapis-project/pkg/api/interfaces"
	"sync"
)

type Kernel struct {
	plugins map[string]i.Plugin
	mu      sync.Mutex
	started bool
}
