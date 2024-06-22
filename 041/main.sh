#!/bin/bash

REPORT_FILE="/tmp/system_report.txt"

# CPU使用率
CPU_USAGE=$(top -bn1 | grep "Cpu(s)" | awk '{print $2 + $4}')

# メモリ使用率
MEM_USAGE=$(free | awk '/Mem/{printf("%.2f"), $3/$2*100}')

# ディスクI/O
DISK_IO=$(iostat -d | awk '/sda/{print $2}')

# レポート生成
echo "System Report $(date)" > "$REPORT_FILE"
echo "CPU Usage: $CPU_USAGE%" >> "$REPORT_FILE"
echo "Memory Usage: $MEM_USAGE%" >> "$REPORT_FILE"
echo "Disk I/O: $DISK_IO KB/s" >> "$REPORT_FILE"
