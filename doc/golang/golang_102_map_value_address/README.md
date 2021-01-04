## case

**map 值地址問題**

請說出下面代碼，執行時為什麼會報錯？

```go
type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"Bach"}}
	m["people"].name = "Goldberg"
}
```

## solution

因爲 map 中的 value 是不可尋址的。在 map 擴容時，會導致 value 進行複製，從而舊的指向 value 的地址失效。

### f1

如要修改結構體的值，那麼應該存入結構體的指針而不是值。代碼見 [函數f1](main.go)。