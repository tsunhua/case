## case

**實現一個阻塞讀且並發安全的 map**

實現一個並發安全的 map，讀取 value 時如 key 不存在，則阻塞直到 key 存在或者超時。

```go
type Smap interface {
	// 寫入 key-value，非阻塞
	Put(key string, val interface{})
	// 根據 key 找 value。如 key 不存在，則則色直到 key 存在或超時
    Get(key string, timeout time.Duration) interface{}
}
```

## solution

1. 使用 RWMutex 實現讀寫的並發安全。
2. 使用一個 map 保存 key-channel 組合，當有讀取操作且 key 不存在時，創建一個 channel，等待 key 存在後再刪除 channel。另一方面，在寫入 key-value 時，檢查是否有 channel，如有則寫入 true 到其中，然後關閉 channel。
3. 超時通過 `time.After(duration)`(返回一個單方向只接收數據的 channel，其中的值是超時的時間點time) 實現。

具體代碼見 [main.go](main.go)