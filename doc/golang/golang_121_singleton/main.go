package main

import (
	"sync"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	once := &Once{}
	for i := 0; i < 5; i++ {
		go once.Do(func() {
			println("Ok")
		})
	}
	time.Sleep(3 * time.Second)
}
