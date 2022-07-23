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
sudo systemctl restart nginx
sudo systemctl restart isuports
