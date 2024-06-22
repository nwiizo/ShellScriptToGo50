#!/bin/bash

LOG_FILE="/var/log/system_update.log"
ADMIN_EMAIL="admin@example.com"

echo "System update started at $(date)" >> "$LOG_FILE"

# システムアップデート
sudo apt update >> "$LOG_FILE" 2>&1
sudo apt upgrade -y >> "$LOG_FILE" 2>&1

# セキュリティアップデートの確認
if sudo apt list --upgradable | grep -q "security"; then
    echo "Important security updates are available." | \
    mail -s "Security Update Alert" "$ADMIN_EMAIL"
fi

echo "System update completed at $(date)" >> "$LOG_FILE"
