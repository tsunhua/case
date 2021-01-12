## case

使用配置文件进行 ssh 登陆服务器

## solution

```bash
mkdir -p ~/.ssh && chmod 700 ~/.ssh
touch ~/.ssh/config
chmod 600 ~/.ssh/config
```

https://linuxize.com/post/using-the-ssh-config-file/

```
Host targaryen
    HostName 192.168.1.10
    User daenerys
    Port 7654
    IdentityFile ~/.ssh/targaryen.key
```

Host tencent
    HostName 139.199.66.91
    User root
    Port 22
    IdentityFile /Users/lshare/Documents/key/tencent/tencent_gz_3.pem