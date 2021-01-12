## case

**iota 賦值問題**

寫出下列代碼打印的內容。

```go
package main
import "fmt"
const (
   a = iota
   b = iota
)
const (
   name = "menglu"
   c    = iota
   d    = iota
)
func main() {
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Println(d)
}
```

## solution

兩個 const 組中的 iota 不是連續的，單個 const 組中，iota 值從0開始遞增，並與常量綁定，無論其是否使用。
因此輸出爲
```
0
1
1
2
```
