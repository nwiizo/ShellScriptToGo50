#!/bin/bash

# config.yaml からデータベース接続情報を読み込む
DB_HOST=$(grep 'host:' config.yaml | awk '{print $2}')
DB_USER=$(grep 'user:' config.yaml | awk '{print $2}')
DB_PASSWORD=$(grep 'password:' config.yaml | awk '{print $2}')
DB_NAME=$(grep 'dbname:' config.yaml | awk '{print $2}')
DB_PORT="5432"

# データベースの作成
# createdb -h $DB_HOST -p $DB_PORT -U $DB_USER -O $DB_USER $DB_NAME

# DDLの実行
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f library_ddl.sql

# DMLの実行
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f library_dml.sql

echo "Database setup is complete."
