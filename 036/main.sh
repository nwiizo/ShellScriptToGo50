#!/bin/bash

directory=".."
num_workers=4

files=$(find "$directory" -type f)

process_file() {
    local file="$1"
    local word_count=$(wc -w < "$file")
    echo "$word_count"
}

worker() {
    local worker_id="$1"
    while IFS= read -r file; do
        word_count=$(process_file "$file")
        echo "$word_count"
    done
}

export -f process_file

total_word_count=0
for ((i=1; i<=num_workers; i++)); do
    word_counts=$(echo "$files" | xargs -I{} -P "$num_workers" bash -c 'process_file "$@"' _ {})
    for word_count in $word_counts; do
        total_word_count=$((total_word_count + word_count))
    done
done

echo "Total word count: $total_word_count"
