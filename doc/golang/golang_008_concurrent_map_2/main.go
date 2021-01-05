package main

import (
	"fmt"
	"sync"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	sync.Mutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}

// true 表示被禁止訪問了
// false 表示未被禁止訪問
func (o *Ban) visit(ip string) (rz bool) {
	o.Lock()
	defer o.Unlock()
	t, ok := o.visitIPs[ip]
	if ok && time.Now().Sub(t) <= 3*time.Minute {
		rz = true
	} else {
		o.visitIPs[ip] = time.Now()
	}
	return
}

func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}(j)
		}

	}
	fmt.Println("success:", success)
}
