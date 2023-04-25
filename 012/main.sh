#!/bin/bash
dir="$1"
total_size=0
for file in "$dir"/*; do
  if [ -f "$file" ]; then
    file_size=$(stat -f%z "$file")
    total_size=$((total_size + file_size))
  fi
done
echo "Total size: $total_size bytes"
