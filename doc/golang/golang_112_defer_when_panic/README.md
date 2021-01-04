## case

**defer是否在panic後被執行的問題**

寫出下面代碼輸出內容

```go
func main() {
    defer_call()
}

func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印後") }()

    panic("觸發異常")
}
```

## solution

panic 不會阻止 defer 函數的執行，defer 按照 LIFO 的順序執行。

```go
打印後
打印中
打印前
```

附：[defer 的實現原理](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/#537-小结)

`defer` 关键字的实现主要依靠编译器和运行时的协作，我们总结一下本节提到的三种机制：

- 堆上分配 · 1.1 ~ 1.12
    - 编译期将 `defer` 关键字被转换 `[runtime.deferproc](https://draveness.me/golang/tree/runtime.deferproc)` 并在调用 `defer` 关键字的函数返回之前插入 `[runtime.deferreturn](https://draveness.me/golang/tree/runtime.deferreturn)`；
    - 运行时调用 `[runtime.deferproc](https://draveness.me/golang/tree/runtime.deferproc)` 会将一个新的 `[runtime._defer](https://draveness.me/golang/tree/runtime._defer)` 结构体追加到当前 Goroutine 的链表头；
    - 运行时调用 `[runtime.deferreturn](https://draveness.me/golang/tree/runtime.deferreturn)` 会从 Goroutine 的链表中取出 `[runtime._defer](https://draveness.me/golang/tree/runtime._defer)`结构并依次执行；
- 栈上分配 · 1.13
    - 当该关键字在函数体中最多执行一次时，编译期间的 `[cmd/compile/internal/gc.state.call](https://draveness.me/golang/tree/cmd/compile/internal/gc.state.call)` 会将结构体分配到栈上并调用 `[runtime.deferprocStack](https://draveness.me/golang/tree/runtime.deferprocStack)`；
- 开放编码 · 1.14 ~ 现在
    - 编译期间判断 `defer` 关键字、`return` 语句的个数确定是否开启开放编码优化；
    - 通过 `deferBits` 和 `[cmd/compile/internal/gc.openDeferInfo](https://draveness.me/golang/tree/cmd/compile/internal/gc.openDeferInfo)` 存储 `defer` 关键字的相关信息；
    - 如果 `defer` 关键字的执行可以在编译期间确定，会在函数返回前直接插入相应的代码，否则会由运行时的 `[runtime.deferreturn](https://draveness.me/golang/tree/runtime.deferreturn)` 处理；