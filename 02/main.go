// +build ignore
package main

import (
	"sync"
)

func main() {

}

// SyncMap is thread safe map struct
type SyncMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// NewSyncMap it to
func NewSyncMap() *SyncMap {
	return &SyncMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

// Get is get value from map
func (m *SyncMap) Get(key interface{}) (v interface{}, ok bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	v, ok = m.data[key]
	return
}

// Set is set value to map[key]
func (m *SyncMap) Set(key interface{}, value interface{}) {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.data[key] = value
}

// Delete is to delete map[key]
func (m *SyncMap) Delete(key interface{}) {
	m.rw.Lock()
	defer m.rw.Unlock()

	delete(m.data, key)
}
