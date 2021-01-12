package main

import "fmt"

const (
	a = iota
	b = iota
	c
)

const d = 33

const (
	e = 33
	f = "menglu"
	g = iota
	h
	i = 44
	j = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(j)
}
