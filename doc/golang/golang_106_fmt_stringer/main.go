package main

import (
	"strings"
)

type People struct {
	Name string
}

func (p *People) String() string {
	sb := &strings.Builder{}
	sb.WriteString("People{")
	sb.WriteString("Name:" + p.Name + ",")
	sb.WriteString("}")
	return sb.String()
}

func main() {
	p := &People{}
	println(p.String())
}
