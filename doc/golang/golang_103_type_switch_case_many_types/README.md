## case

請說出下面代碼存在什麼問題？

```go
type student struct {
	Name string
}

func f1(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name
	}
}
```

## solution

無法確知 msg 爲 *student 或 student 中的哪一個類型，因爲 `msg.Name` 無法執行。當對值v的類型進行 switch case 時，如果 case 中有多個類型，此時值v仍然是 interface{} 。詳見：https://golang.org/ref/spec#Type_switches

代碼見 [main.go](main.go)。

### f1

解決方法很簡單，分開兩句寫即可。
