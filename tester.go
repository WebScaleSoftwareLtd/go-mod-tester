package gomodtester

import "sync"

var (
	m    = map[string]struct{}{}
	lock = sync.RWMutex{}
)

// BlowFuse means that the string specified will return true for the IsBlown method.
func BlowFuse(s string) {
	lock.Lock()
	defer lock.Unlock()
	m[s] = struct{}{}
}

// IsBlown returns true if the string specified has been blown.
func IsBlown(s string) bool {
	lock.RLock()
	defer lock.RUnlock()
	_, ok := m[s]
	return ok
}
