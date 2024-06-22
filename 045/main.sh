#!/bin/bash

LOG_FILE="/var/log/auth.log"
REPORT_FILE="/tmp/failed_logins.txt"

tail -f "$LOG_FILE" | while read line; do
    if echo "$line" | grep -q "Failed password"; then
        echo "$line" >> "$REPORT_FILE"
        echo "Failed login attempt detected: $line"
    fi
done
