package main

import (
	"fmt"
	"sync"
	"time"
)

type Smap interface {
	Put(key string, val interface{})
	Get(key string, timeout time.Duration) interface{}
}

type MSmap struct {
	c    map[string]interface{}
	cham map[string]chan bool
	sync.RWMutex
}

func main() {
	var m Smap = &MSmap{
		c:    make(map[string]interface{}),
		cham: make(map[string]chan bool),
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		m.Put("k", "aa")
		time.Sleep(500 * time.Millisecond)
		m.Put("k", "bb")
	}()

	fmt.Println("1:", m.Get("k", 100*time.Millisecond))
	fmt.Println("2:", m.Get("k", time.Second))
	time.Sleep(600 * time.Millisecond)
	fmt.Println("3:", m.Get("k", time.Second))
}

func (m *MSmap) Put(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()
	m.c[key] = val
	if _, ok := m.cham[key]; ok {
		m.cham[key] <- true
		close(m.cham[key])
	}
}

func (m *MSmap) Get(key string, timeout time.Duration) interface{} {
	if v, ok := m.GetImmediately(key); ok {
		return v
	}

	m.Lock()
	m.cham[key] = make(chan bool)
	m.Unlock()

	defer func() {
		delete(m.cham, key)
	}()

	select {
	case <-m.cham[key]:
		if v, ok := m.GetImmediately(key); ok {
			return v
		}
	case <-time.After(timeout):
	}
	return nil
}

func (m *MSmap) GetImmediately(key string) (v interface{}, ok bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok = m.c[key]
	return
}
