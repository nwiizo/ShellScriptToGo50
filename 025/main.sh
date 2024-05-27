#!/bin/bash

echo "Enter 5 numbers:"
for ((i=1; i<=5; i++)); do
    read -p "Number $i: " number
    numbers+=("$number")
done

sum=0
for num in "${numbers[@]}"; do
    ((sum+=num))
done

average=$(echo "scale=2; $sum / ${#numbers[@]}" | bc)

echo "Sum: $sum"
echo "Average: $average"
