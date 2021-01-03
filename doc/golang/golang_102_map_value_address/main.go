package main

type Student struct {
	name string
}

func main() {
	f1()
}

func f1() {
	m := map[string]*Student{"people": {"Bach"}}
	m["people"].name = "Goldberg"
}
