## case

**atomic 自旋問題**

請說明下面代碼書寫是否正確。

```go
var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v+delta)) {
			break
		}
	}
}
```

## solution

這裏不需要 for 循環。CAS 需不需要自旋要看內存中的值和舊值獲取的時間是否有差異，如果是同一時間獲取的，那麼根本不需要自旋。修改代碼如下：

```go
var value int32

func SetValue(delta int32) {
	atomic.CompareAndSwapInt32(&value, value, value+delta)
}
```