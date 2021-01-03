package main

import (
	"strings"
)

func main() {
	str1 := "abcd123"
	str2 := "dcab321"
	str3 := "dcab329"
	println(f1(str1, str2), f1(str1, str3))
	println(f2(str1, str2), f2(str1, str3))
}

func f1(str1, str2 string) bool {
	if str1 == "" && str2 == "" {
		return true
	}
	if (str1 == "" && str2 != "") || (str1 != "" && str2 == "") {
		return false
	}

	m := make(map[rune]int32)

	for _, r := range []rune(str1) {
		i, _ := m[r]
		m[r] = i + 1
	}

	for _, r := range []rune(str2) {
		i, _ := m[r]
		m[r] = i - 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func f2(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
