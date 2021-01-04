package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
}

func f1() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for a := range ch {
			fmt.Println("a: ", a)
		}
		fmt.Println("close")
	}()
	fmt.Println("ok")
	time.Sleep(time.Second * 3)
}
