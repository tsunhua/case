package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	f1()
	f2()
}

func f1()  {
	s := new(Show)
	s.Param = make(map[string]interface{})
	s.Param["RMB"] = 10000
	fmt.Printf("%v\n", s)
}

func f2()  {
	s := &Show{
		Param: make(map[string]interface{}),
	}
	s.Param["RMB"] = 10000
	fmt.Printf("%v\n", s)
}