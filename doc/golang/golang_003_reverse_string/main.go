package main

func main() {
	str := "abcd1234"
	println(f1(str))
	println(f2(str))
}

func f1(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; j > i; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func f2(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
