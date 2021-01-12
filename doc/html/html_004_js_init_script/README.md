## case

**JS 加載初始化腳本**

## solution

https://blog.csdn.net/DengJH_Business/article/details/32107303

第一种：
```html
<script language=javascript type="text/javascript">
//固定写法
window.onload = function(){
    alert("JS初始化加载");
}
</script>
```


第二种：
```html
<html>
<head>
<script language=javascript type="text/javascript">
//方法名自定义
function init(){
    alert("JS初始化加载");
}
</script>
</head>

<body οnlοad="init()">
</body>

</html>
```

第三种：

```html
<script language=javascript type="text/javascript">
//方法名自定义
init();
 
//方法名要一致
function init(){ 
    alert("JS初始化加载");
}
</script>
```

至于“JavaScript脚本放在哪里”，看到很多资料说：

“

在HTML body部分中的JavaScripts会在页面加载的时候被执行。
在HTML head部分中的JavaScripts会在被调用的时候才执行。

”



但自己试了，放<head></head>，<body></body>，甚至<html></html>，Firefox 30都执行。