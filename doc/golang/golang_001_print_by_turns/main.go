package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	f1()
}

func f1() {
	number, letter := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d%d", i, i+1)
				i = i + 2
				letter <- true
			default:
				break
			}
		}
	}()

	go func(wg *sync.WaitGroup) {
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i+1 > strings.Count(letters, "")-2 {
					wg.Done()
					return
				}
				fmt.Print(letters[i : i+2])
				i = i + 2
				number <- true
			default:
				break
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
