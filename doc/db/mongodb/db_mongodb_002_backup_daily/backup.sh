#!/bin/sh
DATE=`date +%Y%m%d%H%M%S`
SERVER_PATH='/root/apps/myapp/backup'
LOCAL_PATH='/Users/u/Downloads/backup/mydb'

IP='127.0.0.1'
USER='root'
PSW='psw'
DB_NAME='mydb'

## 使用 crontab 設定定時任務，每天早上 6 點和傍晚 6 點執行備份
## sudo crontab -e 
## 0 6,18 * * * /root/apps/repeat/backup.sh -b >> /root/apps/repeat/log/backup.log 2>&1
back(){
    echo "開始備份..."
    mkdir -p $SERVER_PATH
    cd $SERVER_PATH
    mongodump -h 127.0.0.1:27017 --authenticationDatabase admin -u $USER -p $PSW -d $DB_NAME -o $SERVER_PATH
    tar -zcvf ${DB_NAME}"_"${DATE}.tar.gz ${DB_NAME}
    rm -Rf $DB_NAME
    echo "備份完成"
}

## ./backup.sh -s
sync(){
    echo "開始同步..."
    rsync -avz -e ssh tencent:${SERVER_PATH} $LOCAL_PATH
    echo "同步完成"
}

if [ "$1" == '-b' ]; then
    back
elif [ "$1" == '-s' ]; then
    sync
else
    echo "需要參數 -b(備份數據) 或 -s(同步數據)"
fi