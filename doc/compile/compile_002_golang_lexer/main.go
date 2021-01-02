package main

import (
	"fmt"
	"io/ioutil"
	"go/scanner"
	"go/token"
)

const FileName = "hello.go"

func main() {
	src, err := ioutil.ReadFile(FileName)
	if err != nil{
		println("err:",err.Error())
		return
	}

	var fset = token.NewFileSet()
	var file = fset.AddFile(FileName, fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}