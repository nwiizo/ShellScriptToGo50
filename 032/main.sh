#!/bin/bash

url="https://jsonplaceholder.typicode.com/users"
response=$(curl -s "$url")

echo "$response" | jq -r '.[] | "ID: \(.id), Name: \(.name), Email: \(.email)"'
