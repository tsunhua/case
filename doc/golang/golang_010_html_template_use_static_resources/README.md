## case

**Go html 模板使用靜態資源**

## solution

```go
mux.ServeFiles("/css/*filepath", http.Dir("public/css"))
```

```html
<link href="/css/bootstrap.min.css" rel="stylesheet">
```

參考：https://michaelchen.tech/golang-web-programming/css/
