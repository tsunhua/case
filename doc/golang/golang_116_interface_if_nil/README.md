## case

**接口是否爲 nil 的問題**

以下代码打印出来什么内容。

```go
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
```

## solution

打印內容爲：

```
BBBBBBB
```

*Student 定義後沒有初始化，所以 *Student 是 nil，但 *Student 實現了 People 接口，接口不為 nil。