package main

import (
	"sync"
)

type M struct {
	Map  map[string]string
	lock sync.RWMutex
}

func (m *M) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

func (m *M) Get(key string) string {
	return m.Map[key]
}

// TestMap  ...
