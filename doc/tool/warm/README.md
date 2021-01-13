## case

**暖寶寶程序**

天氣很冷，金屬外殼的電腦打字時感覺很凍，不想敲代碼，怎麼辦？

## solution

編寫暖寶寶程序，把 CPU 轉起來，散發熱量來加熱金屬外殼。

```go
func main() {
	for { go func() {}() }
}
```

程序見 [main.go](main.go)