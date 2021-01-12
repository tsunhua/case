## case

**將兩個 div 元素並排**

## solution


### 方法一

參考：https://www.zeusdesign.com.tw/article/3-HTML%20%E5%85%83%E7%B4%A0%E4%B8%A6%E6%8E%92.html

```html
<style type="text/css">
    .clearfix {
        *zoom: 1;
    }
    .clearfix:after {
        content: "";
        display: table;
        line-height: 0;
        clear: both;
    }
    .box {
        border: 1px solid red;
        width: 300px;
    }
    .a, .b {
        display: inline-block;
        width: 100px;
        height: 50px;
        border: 1px solid red;
        float: left;
    }
    .a {
        height: 100px;
    }
</style>
<div class='box clearfix'>
    <div class='a'>This is A</div>
    <div class='b'>This is B</div>
</div>

```

### 方法一

使用 `float:left` 屬性，將兩個 div 同時向左浮動。

```html
<style type="text/css">
#DIV1{
width:200px;　//DIV區塊寬度
line-height:50px;　//DIV區塊高度
padding:20px;　//DIV區塊內距，參閱：CSS padding 內距。
border:2px blue solid;　//DIV區塊邊框，參閱：CSS border 邊框設計。
margin-right:10px;　//靠右外距，參閱：CSS margin 邊界使用介紹範例教學。
float:left;
}
#DIV2{
width:200px;
line-height:50px;
padding:20px;
border:2px green solid;
float:left;
}
</style>
<div id="DIV1">這是並排在左邊的 DIV 區塊</div>
<div id="DIV2">這是並排在右邊的 DIV 區塊</div>
<div style="clear:both;"></div><!--這是用來清除上方的浮動效果-->
```
