#!/bin/bash

LOG_DIR="/var/log/myapp"
MAX_LOGS=5

# 現在のログファイルを圧縮
gzip "$LOG_DIR/app.log"

# 古いログファイルを削除
ls -t "$LOG_DIR"/app.log.gz* | tail -n +$((MAX_LOGS + 1)) | xargs rm -f

# 新しいログファイルを作成
touch "$LOG_DIR/app.log"
