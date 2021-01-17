## case

**Go實現單例模式**

以下是Go中單例的實現代碼，請問是否有問題？

```go
package doublecheck
import (
	"sync"
)
type Once struct {
	m    sync.Mutex
	done uint32
}
func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}
```

## solution

以上代碼雖然沒有正確實現單例，但是 f 只會執行一次。
正確的方式是參考 `sync.Once`。

Go 的內存模型參考：
1. [官方](https://golang.org/ref/mem)
2. [翻譯](https://www.pengrl.com/p/34119/)

Go 的內存模型說明了一個問題，如果要想讓一個協程中對變量的修改對另一個協程可見，那麼需要保證一種 happen before 的關係：
1. 對變量v的寫操作w發生在對變量的讀操作r之前；
2. 其他對變量v的寫操作要麼發生在w之前，要麼發生在r之後。

若要保證 happen before 關係，需要使用同步原語，包括 channel、sync 和 atomic 等。
