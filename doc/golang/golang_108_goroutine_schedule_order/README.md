## case

**Go 程調度順序問題**

請說出下面的代碼存在什麼問題？

```go
type query func(string) string

func exec(name string, vs ...query) string {
    ch := make(chan string)
    fn := func(i int) {
        ch <- vs[i](name)
    }
    for i, _ := range vs {
        go fn(i)
    }
    return <-ch
}

func main() {
    ret := exec("111", func(n string) string {
        return n + "func1"
    }, func(n string) string {
        return n + "func2"
    }, func(n string) string {
        return n + "func3"
    }, func(n string) string {
        return n + "func4"
    })
    fmt.Println(ret)
}
```

## solution

Go 程的調度時間是不確定的。在啟動一系列 go 程後直接 return 可能會導致先入棧的 go 程還未執行，最後可能返回最後入棧的 go 程執行的結果，即 `111func4`。



