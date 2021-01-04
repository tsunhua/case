## case

**方法的指針接收者和值接收者問題**

以下代碼能編譯過去嗎？為什麼？

```go
package main

import (
	"fmt"
)

type People interface {
	Hello() string
}

type Student struct{}

func (stu *Student) Hello() string {
	return "Hello"
}

func main() {
	var peo People = Student{}
	fmt.Println(peo.Hello())
}
```

## solution

編譯失敗，值類型 Student{} 未實現接口People的方法，不能定義為 People類型。

在 golang 語言中，Student 和 *Student 是兩種類型，第一個是表示 Student 本身，第二個是指向 Student 的指針。

### f1

將 peo 賦值為 *Student 類型即可。