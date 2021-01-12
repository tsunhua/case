## case

**變量定義與賦值**

以下代碼遇到了一個麻煩，就是 if true 中的短賦值語句，會導致 items 成爲 if 中的局部變量。而無法給外部的賦值，如取消短變量賦值，則 err 未定義。如何解決？

```go
var items []string
if true {
items, err := getItems()
} else {
items = []string{}
}
println(items)
```

## solution

### f1

在 if 之上定義下 err 即可。