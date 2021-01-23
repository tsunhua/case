## case

**當前用戶無權限刪除數據庫**

## solution

（1）免驗證訪問數據庫

编辑 mongod 配置文件

```bash
vim /etc/mongod.conf
```

然後注釋掉驗證部分：

```
#security:
#  authorization: enabled
```

隨後重啟 mongod 服務。

```bash
sudo systemctl restart mongod
```

（2）刪除數據庫

```bash
mongod
> use my_db
switched to db my_db
> db.dropDatabase
```

（3）重新啟用驗證登錄

步驟（1）的逆過程。