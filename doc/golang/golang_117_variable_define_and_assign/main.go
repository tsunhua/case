package main

func main() {
	f0()
	f1()
}

func f0() {
	var items []string
	if true {
		items, err := getItems()
	} else {
		items = []string{}
	}
	println(items)
}

func getItems() (items []string, err error) {
	return
}

func f1() {
	var items []string
	var err error
	if true {
		items, err = getItems()
	} else {
		items = []string{}
	}
	println(items, err)
}
