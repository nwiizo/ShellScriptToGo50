#!/bin/bash

SOURCE_DIR="/path/to/source"

# サブディレクトリを作成
mkdir -p "$SOURCE_DIR"/{documents,images,videos,others}

# ファイルを整理
find "$SOURCE_DIR" -maxdepth 1 -type f | while read file; do
    case "${file##*.}" in
        pdf|doc|docx|txt) mv "$file" "$SOURCE_DIR/documents/" ;;
        jpg|jpeg|png|gif) mv "$file" "$SOURCE_DIR/images/" ;;
        mp4|avi|mov) mv "$file" "$SOURCE_DIR/videos/" ;;
        *) mv "$file" "$SOURCE_DIR/others/" ;;
    esac
done
