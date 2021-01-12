## case

**爲 WaitGroup 添加 timeout 功能**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    wg := sync.WaitGroup{}
    c := make(chan struct{})
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(num int, close <-chan struct{}) {
            defer wg.Done()
            <-close
            fmt.Println(num)
        }(i, c)
    }

    if WaitTimeout(&wg, time.Second*5) {
        close(c)
        fmt.Println("timeout exit")
    }
    time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    // 要求手写代码
    // 要求sync.WaitGroup支持timeout功能
    // 如果timeout到了超时时间返回true
    // 如果WaitGroup自然结束返回false
}
```

## solution

wg.Wait 函數是阻塞的，因此 WaitTimeout 函數也必須是阻塞的。
timeout 先到就返回true，自然結束先到就返回false。使用一個不帶緩衝的 chan 可以實現。


代碼請看：[main.go](main.go)