## case

**map 並發問題**

下面的代碼有什麼問題?

```go
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
```

## solution

在執行 Get方法時可能被panic。

雖然有使用sync.Mutex做寫鎖，但是map是併發讀寫不安全的。map屬於引用類型，併發讀寫時多個協程見是通過指針訪問同一個地址，即訪問共享變量，此時同時讀寫資源存在競爭關係。會報錯誤信息:「fatal error: concurrent map read and map write」。

因此，在 Get 中也需要加鎖，因為這裡只是讀，建議使用讀寫鎖 sync.RWMutex。

修改後的代碼見 [main.go](main.go)