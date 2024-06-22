#!/bin/bash

DB_NAME="mydb"
BACKUP_DIR="/path/to/backups"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/${DB_NAME}_${DATE}.sql"

# バックアップ作成
pg_dump "$DB_NAME" > "$BACKUP_FILE"

if [ $? -eq 0 ]; then
    echo "Backup created successfully: $BACKUP_FILE"
else
    echo "Backup failed"
    exit 1
fi

# 復元関数
restore_db() {
    local backup_file=$1
    if [ -f "$backup_file" ]; then
        psql "$DB_NAME" < "$backup_file"
        echo "Database restored from $backup_file"
    else
        echo "Backup file not found: $backup_file"
    fi
}

# 使用例: ./script.sh restore /path/to/backups/mydb_20230101_120000.sql
if [ "$1" = "restore" ] && [ -n "$2" ]; then
    restore_db "$2"
fi
