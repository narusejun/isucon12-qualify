#!/bin/bash -eux

# 各種設定ファイルのコピー
sudo cp -f env.sh /home/isucon/env.sh

# ミドルウェア・Appの再起動
sudo systemctl restart mysql
sudo systemctl stop nginx
sudo systemctl disable nginx
sudo systemctl stop isuports
sudo systemctl disable isuports
