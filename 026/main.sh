#!/bin/bash

read -p "Enter a string: " input

reverse=$(echo "$input" | rev)

if [[ "$input" == "$reverse" ]]; then
    echo "'$input' is a palindrome."
else
    echo "'$input' is not a palindrome."
fi
