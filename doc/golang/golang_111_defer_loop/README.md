## case

**在loop中使用defer的問題**

下面的程序運行後為什麼會爆異常

```go
type Project struct{}

func (p *Project) deferError() {
    if err := recover(); err != nil {
        fmt.Println("recover: ", err)
    }
}

func (p *Project) exec(msgchan chan interface{}) {
    for msg := range msgchan {
        m := msg.(int)
        fmt.Println("msg: ", m)
    }
}

func (p *Project) run(msgchan chan interface{}) {
    for {
        defer p.deferError()
        go p.exec(msgchan)
        time.Sleep(time.Second * 2)
    }
}

func (p *Project) Main() {
    a := make(chan interface{}, 100)
    go p.run(a)
    go func() {
        for {
            a <- "1"
            time.Sleep(time.Second)
        }
    }()
    time.Sleep(time.Second * 100000000000000)
}

func main() {
    p := new(Project)
    p.Main()
}
```

## solution

1. 100000000000000 超過 time.Duration (int64)。
2. 字符串無法強制轉爲 int 值，需要使用 `strconv.Atoi` 函數。
3. `defer p.deferError()` 不能在循環中調用。

    ```go
    func deferLoop() {
        filenames := []string{"1.txt", "2.txt"}

        for _, filename := range filenames {
            fp, err := os.Open(filename)

            if err != nil {
                log.Printf("open file %s failed, err:%v", filename, err)
                continue
            }

            // Possible resource leak, 'defer' is called in a for loop.
            // 這句語句會造成循環結束後才開始回收資源，而不是循環一次就回收一次資源
            defer fp.Close()

            // 使用fp(file pointer)
        }
    }
    ```

4. `defer p.deferError()` 可以捕獲 run 方法中的 panic，但無法捕獲 exec 方法中的 panic，因爲只能捕獲當前協程的 panic。