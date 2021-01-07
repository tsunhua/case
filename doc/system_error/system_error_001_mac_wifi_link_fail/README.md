## case

**Mac 無法連接上 WiFi**

## solution

1. 打開「設置-網絡-位置」，更改為自動。
2. 打開「設置-網絡-高級-DNS」，設置為 114.114.114.114 和 8.8.8.8。
3. 打開 「Finder-前往-電腦-Macintosh-HD-資源庫-Perferences-SystemConfiguration」，刪除文件 `com.apple.wifi.message-tracer`。
4. 關機重啟。