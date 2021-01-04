## case

**Go 程調度時間不確定問題**

請找出下面代碼的問題所在。

```go
func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
```

## solution

Go 程的調度時間是不確定的，在開啟 go 程執行任務的代碼後直接 close 信道會導致寫數據進入信道失敗，從而引發 panic: send on closed channel。正確的做法是在發送完畢後關閉信道，還有可以使用 for range 簡化代碼。代碼見 [函數f1](main.go)。