#!/bin/bash

target_dir="$1"

for file in "$target_dir"/*; do
  if [ -f "$file" ]; then
    size=$(stat -f%z "$file")
    echo "File: $file, Size: $size bytes"
  fi
done
