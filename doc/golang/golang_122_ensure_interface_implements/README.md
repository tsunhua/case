## case

**保證接口一定被自己實現**

有以下 Error 接口，請保證該接口有且只有被err結構體實現。

```go
package main

type Error interface {
	WithData(data interface{}) Error
	WithID(id string) Error
}

type err struct {
	Code int         `json:"code"`         // 業務編碼
	Msg  string      `json:"msg"`          // 錯誤描述
	Data interface{} `json:"data"`         // 成功時返回的數據
	ID   string      `json:"id,omitempty"` // 當前請求的唯一ID，便於問題定位，忽略也可以
}

func (e *err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e *err) WithID(id string) Error {
	e.ID = id
	return e
}
```

## solution

如果 *err 可以賦值給接口 Error，則 *err 實現了 Error；如果接口中有包外不可見的接口方法，那麼該接口只能在該包內被實現。
具體實現見：[main.go](main.go)
參考：[Go - 统一定义 API 错误码](https://mp.weixin.qq.com/s/C4kjPcYNtq32rwGR2CyYug)
