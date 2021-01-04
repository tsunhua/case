## case

**交替打印數字和字母**

## solution

### f1

使用兩個 channel，兩個 go 程。go 程A負責打印數字，然後通過 channel 通知 go 程B去打印字母；go 程B反之。再使用 WaitGroup 當所有打印完畢後退出程序。代碼見 [函數f1](main.go)。