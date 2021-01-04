## case

**判斷兩個字符串是否排序後相同**

## solution

### f1

1. str1 中的每個字符都可以在 str2 中找到；
2. str1 中每個字符的數量與 str2 中的一樣。

可以設定一個 map(字符-數量)，然後 str1 每次發現一個字符就往 map +1，而 str2 相反(-1)。最後遍歷 map 看裏面的值是否都爲 0 即可。代碼見 [函數f1](main.go)。

### f2

這裡還是使用golang內置方法 strings.Count 來判斷字符是否一致。代碼見 [函數f2](main.go)。