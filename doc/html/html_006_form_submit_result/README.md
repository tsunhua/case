## case

**獲取表單提交結果**

## solution

```html
<script>
    const form = $('#loginForm')
    form.submit(function (event) {
        $.ajax({
            url: form.attr('action'),
            type: 'POST',
            data: form.serialize(),
            success: function () {
                alert("登入成功，將自動跳轉主頁")
                // redirect to add page
                window.location.href = "/add"
            },
            error: function () {
                alert("郵箱或密碼錯誤，請檢查後重試")
            }
        });
        event.preventDefault()
        this.reset()
    });
</script>
```