## case

**兩個結構體是否相等問題**

寫出下面代碼的輸出結果。

```go
package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
```

## solution

輸出內容爲：
```
false
true
```

指針類型比較的是指針地址，非指針類型比較的是每個屬性的值。
