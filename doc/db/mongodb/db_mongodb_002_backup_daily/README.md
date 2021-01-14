## case

**MongoDB 每日備份**

在自己電腦上同步服務器中部署的 MongoDB 中的數據。

## solution

思路：
1. 使用 mongodump 配合 bash 腳本每日備份數據到固定的文件夾。
2. 使用 rsync 配合 bash 腳本每日增量同步服務器上的文件夾。

注意：rsync 必須是服務器和電腦上都安裝才能使用。

代碼參見：[backup.sh](backup.sh)

