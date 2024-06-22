#!/bin/bash

THRESHOLD=90
EMAIL="admin@example.com"

USAGE=$(df -h / | awk 'NR==2 {print $5}' | sed 's/%//')

if [ "$USAGE" -gt "$THRESHOLD" ]; then
    echo "Disk usage is at $USAGE%, which exceeds the threshold of $THRESHOLD%." | \
    mail -s "Disk Usage Alert" "$EMAIL"
fi
