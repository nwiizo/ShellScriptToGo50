#!/bin/bash
filename="sample.txt"

line_number=1
while IFS= read -r line; do
    echo "Line $line_number: ${#line} characters"
    ((line_number++))
done < "$filename"
