## case

已知有 Go 文件 [hello.go](hello.go)，請使用 Go 語言自帶的詞法分析器進行分析。

## solution

使用 `io/ioutil` 讀取文件，使用 `go/scanner` 掃描並解析成 Token 序列，代碼見 [main.go](main.go)

執行 main.go

```bash
go run main.go
```

輸出結果爲：

```
hello.go:1:1    package "package"
hello.go:1:9    IDENT   "main"
hello.go:1:13   ;       "\n"
hello.go:3:1    import  "import"
hello.go:3:8    (       ""
hello.go:4:2    STRING  "\"fmt\""
hello.go:4:7    ;       "\n"
hello.go:5:1    )       ""
hello.go:5:2    ;       "\n"
hello.go:7:1    func    "func"
hello.go:7:6    IDENT   "main"
hello.go:7:10   (       ""
hello.go:7:11   )       ""
hello.go:7:13   {       ""
hello.go:8:2    IDENT   "fmt"
hello.go:8:5    .       ""
hello.go:8:6    IDENT   "Println"
hello.go:8:13   (       ""
hello.go:8:14   STRING  "\"Hello\""
hello.go:8:21   )       ""
hello.go:8:22   ;       "\n"
hello.go:9:1    }       ""
hello.go:9:2    ;       "\n"
```