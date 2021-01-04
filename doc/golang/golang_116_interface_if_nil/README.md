## case

**接口是否爲 nil 的問題**

以下代碼打印出來什麼內容。

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

*Student 定義後沒有初始化，所以 *Student 是 nil，但 *Student 實現了 People 接口，接口不爲 nil。
這點跟 Java 很不一樣，Java 中的話很好理解，此時就都是 null。