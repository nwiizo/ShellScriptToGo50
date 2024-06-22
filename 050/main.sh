#!/bin/bash

LOG_FILE="/var/log/container_monitor.log"

log_message() {
    echo "$(date): $1" >> "$LOG_FILE"
}

restart_container() {
    local container_id=$1
    log_message "Restarting container $container_id"
    docker start "$container_id"
}

log_resource_usage() {
    local container_id=$1
    local cpu_usage=$(docker stats --no-stream --format "{{.CPUPerc}}" "$container_id")
    local mem_usage=$(docker stats --no-stream --format "{{.MemUsage}}" "$container_id")
    log_message "Container $container_id - CPU: $cpu_usage, Memory: $mem_usage"
}

while true; do
    for container_id in $(docker ps -q); do
        if [ "$(docker inspect -f '{{.State.Running}}' "$container_id")" = "false" ]; then
            restart_container "$container_id"
        else
            log_resource_usage "$container_id"
        fi
    done
    sleep 60
done
