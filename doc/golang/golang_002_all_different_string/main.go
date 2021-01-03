package main

import (
	"strings"
)

func main() {
	str1 := "adbcca1992"
	str2 := "udkz;e01jab"
	println(f1(str1), f1(str2))
}

func f1(str string) bool {
	if len(str) > 3000 {
		return false
	}
	for k, v := range str {
		if v > 127 { // ascii 0～127 是鍵盤上可以敲出來的
			return false
		}
		if strings.Index(str, string(v)) != k {
			return false
		}
	}
	return true
}
