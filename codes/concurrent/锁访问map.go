package main

import (
	"fmt"
	"sync"
	"time"
)

// M
type M struct {
	Map  map[interface{}]interface{}
	lock *sync.RWMutex
}

// Set ...
func (m *M) Set(key interface{}, value interface{}) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key interface{}) interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

func main() {
	Map := &M{
		make(map[interface{}]interface{}),
		new(sync.RWMutex),
	}

	Map.Set(2, "2323")
	Map.Set("hello", 343)
	fmt.Println(Map.Get(2))
	fmt.Println(Map.Get("hello"))

	go func() {
		for {
			fmt.Println(Map.Get(2))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {

			Map.Set(2, i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(20 * time.Second)
}
