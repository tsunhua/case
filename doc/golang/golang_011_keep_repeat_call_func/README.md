## case

**每隔1秒調用函數並保證函數不退出**

```go
package main

func main() {
    go func() {
        // 1 在这里需要你写算法
        // 2 要求每秒钟调用一次proc函数
        // 3 要求程序不能退出
    }()

    select {}
}

func proc() {
    panic("ok")
}
```

## solution

調用 time.Tick 方法可以獲得一個只接收的信道，通過 for select 此信道可以間隔 1 秒執行 proc 函數。要想程序不退出，則發生 panic 的函數必須跟 tick 的函數不是同一個。因爲發生panic之後所在的函數不會繼續了。

代碼請看：[main.go](main.go)