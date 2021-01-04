package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})
	Rd(key string, timeout time.Duration) interface{}
}

type MMap struct {
	c    map[string]interface{}
	cham map[string]chan bool
	sync.RWMutex
}

func main() {
	m := MMap{
		c:    make(map[string]interface{}),
		cham: make(map[string]chan bool),
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		m.Out("k", "aa")
		time.Sleep(500 * time.Millisecond)
		m.Out("k", "bb")
	}()

	fmt.Println("1:", m.Rd("k", 100*time.Millisecond))
	fmt.Println("2:", m.Rd("k", time.Second))
	time.Sleep(600 * time.Millisecond)
	fmt.Println("3:", m.Rd("k", time.Second))
}

func (m *MMap) Out(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()
	m.c[key] = val
	if _, ok := m.cham[key]; ok {
		m.cham[key] <- true
		close(m.cham[key])
	}
}

func (m *MMap) Rd(key string, timeout time.Duration) interface{} {
	m.RLock()
	if v, ok := m.c[key]; ok {
		m.RUnlock()
		return v
	}
	m.RUnlock()

	m.Lock()
	m.cham[key] = make(chan bool)
	m.Unlock()

	defer func() {
		delete(m.cham, key)
	}()

	select {
	case <-m.cham[key]:
		m.RLock()
		if v, ok := m.c[key]; ok {
			m.RUnlock()
			return v
		}
		m.RUnlock()
	case <-time.After(timeout):
	}
	return nil
}
