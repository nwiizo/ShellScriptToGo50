#!/bin/sh
file="example.txt"
search_string="about.html"

grep -n "$search_string" "$file"
