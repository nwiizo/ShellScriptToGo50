#!/bin/bash

# config.yaml からデータベース接続情報を読み込む
DB_HOST=$(grep 'host:' config.yaml | awk '{print $2}')
DB_USER=$(grep 'user:' config.yaml | awk '{print $2}')
DB_PASSWORD=$(grep 'password:' config.yaml | awk '{print $2}')
DB_NAME=$(grep 'dbname:' config.yaml | awk '{print $2}')
DB_PORT="5432"

# データベースを削除
dropdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME

# データベースを再作成
createdb -h $DB_HOST -p $DB_PORT -U $DB_USER -O $DB_USER $DB_NAME
