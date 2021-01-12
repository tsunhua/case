
## 配置账号密码

https://blog.fundebug.com/2019/01/21/how-to-protect-mongodb/

账号为”myUserAdmin”，密码为”abc123”。

```
use admin
db.createUser(
  {
    user: "myUserAdmin",
    pwd: "abc123",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)
```

修改MongoDB的配置文件

vim /etc/mongod.conf
将security.authorization设为”enabled”：

security:
  authorization: enabled
重启MongoDB

sudo service mongod restart
连接mongodb

再次连接mongodb时，则需要指定账号与密码。

mongo -u "myUserAdmin" -p "abc123" --authenticationDatabase "admin"
