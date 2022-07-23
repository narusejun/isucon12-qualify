#!/bin/bash -eux

# 各種設定ファイルのコピー
sudo cp -f env.sh /home/isucon/env.sh
sudo cp -f etc/mysql/mysql.conf.d/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf
sudo cp -f etc/nginx/nginx.conf /etc/nginx/nginx.conf
sudo cp -f etc/nginx/sites-available/isuports.conf /etc/nginx/sites-available/isuports.conf
sudo nginx -t

# アプリケーションのビルド
cd /home/isucon/webapp/go
make

# ミドルウェア・Appの再起動
sudo systemctl restart mysql
sudo systemctl reload nginx
sudo systemctl restart isuports

# slow query logの有効化
QUERY="
set global slow_query_log_file = '/var/log/mysql/mysql-slow.log';
set global long_query_time = 0;
set global slow_query_log = ON;
"
echo $QUERY | mysql -uroot -proot

# log permission
sudo chmod 777 /var/log/nginx /var/log/nginx/*
sudo chmod 777 /var/log/mysql /var/log/mysql/*
