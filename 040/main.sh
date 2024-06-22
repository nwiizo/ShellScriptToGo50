#!/bin/bash

SOURCE_DIR="/path/to/source"
BACKUP_DIR="/path/to/backup"
DATE=$(date +%Y%m%d)

# 最新の完全バックアップを見つける
LAST_FULL=$(ls -t "$BACKUP_DIR"/full_* | head -n1)

# 増分バックアップを作成
tar -czf "$BACKUP_DIR/incr_$DATE.tar.gz" \
    -g "$BACKUP_DIR/snapshot" \
    -C "$(dirname "$SOURCE_DIR")" "$(basename "$SOURCE_DIR")"

# スナップショットファイルが存在しない場合、これは完全バックアップとなる
if [ ! -f "$BACKUP_DIR/snapshot" ]; then
    mv "$BACKUP_DIR/incr_$DATE.tar.gz" "$BACKUP_DIR/full_$DATE.tar.gz"
fi
