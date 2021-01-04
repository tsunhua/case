## case

**序列化未導出的成員變量**

寫出打印的結果

```go
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
```

## solution

輸出結果爲：
```
people:  {}
```

結構體 People 中的屬性 name 首字母小寫，因爲是私有的，無法進行 JSON 序列化和反序列化。json 包中的 Marshal 方法注釋裏講到：

Struct values encode as JSON objects. Each exported struct field becomes a member of the object, using the field name as the object key, unless the field is omitted for one of the reasons given below.// Struct 值會被編碼爲一個 JSON 對象。每個導出的結構體字段會成為該對象的成員，使用字段名作爲對象的 key，除非該字段因爲以下原因之一被忽略。

### f1

將 People 中的 name 字段改爲 Name 即可。代碼見 [函數f1](main.go)。