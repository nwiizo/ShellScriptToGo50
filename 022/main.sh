#!/bin/bash
watch_dir="."

inotifywait -m "$watch_dir" -e create --format '%w%f' |
while read file; do
    echo "New file created: $file"
done
