#!/bin/sh
dir_count=$(find .. -type d | wc -l)
file_count=$(find .. -type f | wc -l)
echo "Directories: $dir_count, Files: $file_count"
