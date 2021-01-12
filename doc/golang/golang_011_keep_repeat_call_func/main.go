package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
	//f2()
}

func f1() {
	go func() {
		c := time.Tick(1 * time.Second)
		for {
			select {
			case <-c:
				func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Printf("%v", err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {} // 這個有什麼作用？Block current goroutine，注釋這句，程序直接退出，不會運行到 Go程中的代碼
	// 等價於
	//var c chan struct{}
	//c <- struct{}{}
}

func f2() {
	go func() {
		defer func() {
			if err := recover(); err != nil { // recover 之後，發生 panic 的函數不會繼續運行。
				fmt.Printf("%v", err)
			}
		}()
		c := time.Tick(1 * time.Second)
		for {
			select {
			case <-c:
				proc()
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
