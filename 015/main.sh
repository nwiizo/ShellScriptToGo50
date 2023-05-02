#!/bin/bash

target_dir="$1"
extension_file="extension_count.txt"

list_files() {
  dir="$1"
  find "$dir" -type f
}

for file in $(list_files "$target_dir"); do
  ext="${file##*.}"
  if [ "$ext" == "$file" ]; then
    ext="other"
  fi
  echo "$ext" >> "$extension_file"
done

sort "$extension_file" | uniq -c | while read count ext; do
  echo "Extension: $ext, Count: $count"
done

rm "$extension_file"
