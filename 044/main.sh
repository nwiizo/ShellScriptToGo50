#!/bin/bash

HOST="example.com"
INTERVAL=60
EMAIL="admin@example.com"

while true; do
    if ! ping -c 1 "$HOST" &> /dev/null; then
        echo "Connection to $HOST failed at $(date)" | \
        mail -s "Network Connection Alert" "$EMAIL"
    fi
    sleep "$INTERVAL"
done
