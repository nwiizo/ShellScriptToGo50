#!/bin/bash

str1="listen"
str2="silent"

sorted1=$(echo "$str1" | tr '[:upper:]' '[:lower:]' | grep -o . | sort | tr -d '\n')
sorted2=$(echo "$str2" | tr '[:upper:]' '[:lower:]' | grep -o . | sort | tr -d '\n')

if [[ "$sorted1" == "$sorted2" ]]; then
    echo "'$str1' and '$str2' are anagrams."
else
    echo "'$str1' and '$str2' are not anagrams."
fi
