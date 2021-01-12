## case

**提交表單後清除輸入內容**

## solution

### 0

```html
<script>
    window.onload = function () {
        const addRepeatForm = document.forms.namedItem("addRepeatForm")
        addRepeatForm.addEventListener('submit', function (event) {
            // 沒有下面這一行將不會把input的值submit出去
            event.preventDefault()
            addRepeatForm.submit()
            addRepeatForm.reset()
        });
    }
</script>
```

### 1
```html
<script type="text/javascript">
    function clearThis(target){
        if(target.value=='exemplo@exemplo.com'){
        target.value= "";}
    }
    </script>
```
https://stackoverflow.com/questions/17237772/html-how-to-clear-input-using-javascript

```html
<button onclick="document.getElementById('myInput').value = ''">Clear input field</button>
<input type="text" value="Blabla" id="myInput">
```
https://www.w3schools.com/howto/howto_html_clear_input.asp


### 2

```html
<input type="button" value="Submit" id="btnsubmit" onclick="submitForm()">
<script>
    function submitForm() {
        // Get the first form with the name
        // Usually the form name is not repeated
        // but duplicate names are possible in HTML
        // Therefore to work around the issue, enforce the correct index
        var frm = document.getElementsByName('contact-form')[0];
        frm.submit(); // Submit the form
        frm.reset();  // Reset all form data
        return false; // Prevent page refresh
    }
</script>
```
https://stackoverflow.com/questions/14589193/clearing-my-form-inputs-after-submission/14589251