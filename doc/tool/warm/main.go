package main

func main() {
	for { go func() {}() }
}
