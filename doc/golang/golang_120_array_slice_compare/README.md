## case

**數組或切片比較**

指出下面代碼的問題。

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println([...]string{"1"} == [...]string{"1"})
    fmt.Println([]string{"1"} == []string{"1"})
}
```

## solution

切片不能相互比較。
數組只能與相同元素類型和數量的進行比較。

