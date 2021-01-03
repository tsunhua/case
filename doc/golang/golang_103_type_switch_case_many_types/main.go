package main

func main() {
	f1(&student{Name: "Bach"})
	f1(student{Name: "Bach"})
}

type student struct {
	Name string
}

func f1(v interface{}) {
	switch msg := v.(type) {
	case *student:
		println(msg.Name)
	case student:
		println(msg.Name)
	}
}
