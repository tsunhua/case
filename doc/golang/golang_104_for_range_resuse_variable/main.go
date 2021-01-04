package main

type student struct {
	Name string
	Age  int
}

func main() {
	// f1()
	f2()
}

func f1()  {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		tmp := stu // 使用中間的臨時變量
		m[stu.Name] = &tmp
	}
}

func f2()  {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i := range stus {
		stu := stus[i] // 從 stus 中取每個 stu 的地址
		m[stu.Name] = &stu
	}
}