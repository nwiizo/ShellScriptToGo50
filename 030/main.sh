#!/bin/bash

url="https://blog.3-shake.com/"

start_time=$(date +%s.%N)
response=$(curl -s -o /dev/null -w "%{http_code}" "$url")
sleep 1
end_time=$(date +%s.%N)

response_time=$(echo "$end_time - $start_time" | bc)
response_time_ms=$(printf "%.3f" "$response_time")

echo "URL: $url"
echo "Status Code: $response"
echo "Response Time: $response_time_ms ms"
