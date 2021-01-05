## case

**高並發下的鎖與map的讀寫**

在一個高併發的web服務器中，要限制IP的頻繁訪問。現模擬100個IP同時併發訪問服務器，每個IP要重復訪問1000次。
每個IP三分鐘之內只能訪問一次。修改以下代碼完成該過程，要求能成功輸出 success:100

```go
package main

import (
	"fmt"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()
		}

	}
	fmt.Println("success:", success)
}
```

## solution

代碼中的問題點有：
1. 並發環境下 map 沒有加鎖，會導致並發錯誤；
2. go 程訪問外部變量，但期望是創建 go 程時的值，這不現實；
3. 未通過當前時間和已經保存的時間點進行比較，保證同個 IP 的訪問間隔不少於 3 分鐘。

修改後的代碼見：[main.go](main.go)
