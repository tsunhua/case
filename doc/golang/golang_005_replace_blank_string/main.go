package main

import (
	"errors"
	"strings"
	"unicode"
)

func main() {
	str := "ab cd  ef g"
	println(f1(str))
	println(f2(str))
}

func f1(origin string) string {
	runes := []rune(origin)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			newRunes = append(newRunes, rune(r))
		} else if r == ' ' {
			newRunes = append(newRunes, '%', '2', '0')
		}
	}
	return string(newRunes)
}

func f2(origin string) (string, error) {
	if len([]rune(origin)) > 1000 {
		return origin, errors.New("len > 1000")
	}
	for _, s := range origin {
		if s != ' ' && !unicode.IsLetter(s) {
			return origin, errors.New("include invalid character")
		}
	}
	return strings.Replace(origin, " ", "%20", -1), nil
}
