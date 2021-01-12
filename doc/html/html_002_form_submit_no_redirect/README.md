## case

**提交表單而不跳轉頁面**

## solution

將 form 中 的 target 屬性設置爲一個隱藏的 iframe 即可。

```html
<form id="addRepeatForm" method="post" action="/action/add" target="hide_iframe">
    <div class="mb-3">
        <label for="contentInput" class="form-label">內容</label>
        <textarea class="form-control" id="contentInput" rows="6" name="content" required></textarea>
    </div>
    <div>
        <button type="submit" class="btn btn-primary d-inline-block">添加</button>
    </div>
</form>
<iframe name="hide_iframe" style="display:none;"></iframe>
```

