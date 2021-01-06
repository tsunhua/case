## case

使用 go template 構建一個站點。

## solution

Go 標準庫提供了兩個模板輸出庫，一個是 text/template，一個是 html/template。
這裡採用 gin 作爲 web 引擎，採用 html/template 輸出 html 內容。

### 起步

go html template 分兩部分，一個是模板文件，一個是解析函數。
模板文件是一個 html 文件，使用 `{{ }}` 包裹動作，其當前上下文是是 `.`，代表傳入模板的數據。
常用的解析函數有：
1. New // 創建一個模板
2. Funcs // 設定可供模板調用的函數
3. ParseFiles // 解析模板文件
4. Must // 包裹返回值爲 (*Template, error) 的函數，當 error 發生時拋出 panic
5. Execute // 執行解析

代碼請看 [函數 hello](main.go)。

### 若...則...

模板中的 if else 格式如下：

```go
{{if expr}}
{{else if expr}}
{{else}}
{{end}}
```
比較大小不能用 >,>=,<,<=,= 等符號，只能用 gt,ge,lt,le,eq 等內置函數。
&,| 也不能用，只能用函數 and,or。

可用`()` 包括函數名及參數。

代碼請看 [函數 ifelse](main.go)。

### 遍歷

模板中的 rang 可以遍歷數組、Map 等。

代碼請看 [函數 rangef](main.go)。

### 比較大小

比較大小不能用 >,>=,<,<=,= 等符號，只能用 gt,ge,lt,le,eq 等內置函數。
&,| 也不能用，只能用函數 and,or。

可用`()` 包括函數名及參數。

代碼請看 [函數 compare](main.go)。

### 調用結構體方法

注意：只能在 gohtml 中調用接收者爲結構體的方法，不能調用接收者爲指針類型的方法。

代碼請看 [函數 callmethod](main.go)。

### 調用自定義的函數

通過 template.Funcs 函數，傳入一個 template.FuncMap。

注意：Funcs 必須在 ParseFiles 之前，否則在 ParseFiles 時將會拋出函數未定義錯誤。

代碼請看 [函數 callfunction](main.go)。

### 使用第三方定義的函數庫 [sprig](https://github.com/Masterminds/sprig)

所有支持的函數見：[http://masterminds.github.io/sprig/](http://masterminds.github.io/sprig/)

代碼請看 [函數 usesprig](main.go)。

### 嵌套模板

[函數 nestedTemplate](main.go) 展示了一個名爲 base 的模板嵌套了 layout 和 footer 模板，並進行了值傳遞。

### 參考文檔

1. [[译]Golang template 小抄](https://colobu.com/2019/11/05/Golang-Templates-Cheatsheet/)
2. [Go Web 编程之 模板（一）](https://darjun.github.io/2019/12/31/goweb/template1/)
3. [Go Web 编程之 模板（二）](https://darjun.github.io/2020/01/09/goweb/template2/)
3. [html/template - golang.org](https://golang.org/pkg/html/template/)
