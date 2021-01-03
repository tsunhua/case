## case

嵌套 struct 的初始化問題。

下面代碼能運行嗎？為什麼？

```go
type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
```

## solution

不能運行，因爲 Param 會被初始化爲 nil，無法對其進行賦值操作。

代碼見 [main.go](main.go)。

### f1

在使用前使用 make 函數進行 map 初始化。

### f2

直接在賦值時進行 map 初始化。
