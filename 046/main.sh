#!/bin/bash

REPORT_FILE="/tmp/performance_report.txt"

echo "Performance Report $(date)" > "$REPORT_FILE"
echo "Top 5 CPU consuming processes:" >> "$REPORT_FILE"
ps aux --sort=-%cpu | head -6 >> "$REPORT_FILE"

echo "" >> "$REPORT_FILE"
echo "Top 5 Memory consuming processes:" >> "$REPORT_FILE"
ps aux --sort=-%mem | head -6 >> "$REPORT_FILE"

echo "Performance report generated at $REPORT_FILE"
