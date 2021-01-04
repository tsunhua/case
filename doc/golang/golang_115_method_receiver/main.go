package main

import (
	"fmt"
)

type People interface {
	Hello() string
}

type Student struct{}

func (stu *Student) Hello() string {
	return "Hello"
}

func main() {
	var peo People = &Student{}
	fmt.Println(peo.Hello())
}
