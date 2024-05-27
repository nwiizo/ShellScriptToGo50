#!/bin/bash
command="ls -l"
timestamp=$(date +%Y%m%d%H%M%S)
output_file="output_${timestamp}.txt"

$command > "$output_file"

echo "Command output saved to $output_file"

echo "Saved files:"
ls -1 output_*.txt
