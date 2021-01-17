package main

// 將 nil 轉爲 *err 再賦值給一個省略變量名的類型爲 Error 的變量
// 保證 *err 一定實現了 Error 接口，否則此處會在編譯時報錯
var _ Error = (*err)(nil)

type Error interface {
	// 私有接口方法，保證該接口只能在包內實現
	i()
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

func (e *err) i(){}
