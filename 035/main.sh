#!/bin/bash

numbers=(10 5 8 12 3 7)

min=${numbers[0]}
max=${numbers[0]}

for num in "${numbers[@]}"; do
    if ((num < min)); then
        min=$num
    fi
    if ((num > max)); then
        max=$num
    fi
done

echo "Minimum: $min, Maximum: $max"
