## case

**for range 重用變量問題**

以下代碼有什麼問題，說明原因。

```go
type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}
```

## solution

for range 循環中進行優化，stu 變量會被複用，每次循環到將值賦到 stu 中，最後 map 中保存的其實都是最後一個 stu。

### f1

使用臨時變量保存 stu，然後再賦值到 map 中。代碼見 [函數f1](main.go)。

### f2

從 stus 中按索引取出 stu，再賦值到 map 中。代碼見 [函數f2](main.go)。