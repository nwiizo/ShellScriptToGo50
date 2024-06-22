#!/bin/bash

LOG_FILE="/var/log/apache2/access.log"

echo "Top 5 IP addresses:"
awk '{print $1}' "$LOG_FILE" | sort | uniq -c | sort -rn | head -5

echo "Top 5 requested URLs:"
awk '{print $7}' "$LOG_FILE" | sort | uniq -c | sort -rn | head -5
