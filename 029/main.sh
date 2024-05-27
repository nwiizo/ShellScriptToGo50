#!/bin/bash

ip_address="8.8.8.8"

ping_output=$(ping -c 4 "$ip_address")

packets_sent=$(echo "$ping_output" | grep -oP '(?<=transmitted, )[0-9]+')
packets_received=$(echo "$ping_output" | grep -oP '(?<=received, )[0-9]+')
packets_lost=$((packets_sent - packets_received))
loss_percentage=$(echo "scale=0; $packets_lost / $packets_sent * 100" | bc)

min_rtt=$(echo "$ping_output" | grep -oP '(?<=min/avg/max/mdev = )[0-9]+\.[0-9]+')
avg_rtt=$(echo "$ping_output" | grep -oP '(?<=min/avg/max/mdev = [0-9]+\.[0-9]+/)[0-9]+\.[0-9]+')
max_rtt=$(echo "$ping_output" | grep -oP '(?<=min/avg/max/mdev = [0-9]+\.[0-9]+/[0-9]+\.[0-9]+/)[0-9]+\.[0-9]+')

echo "Ping statistics for $ip_address:"
echo "  Packets: Sent = $packets_sent, Received = $packets_received, Lost = $packets_lost ($loss_percentage% loss),"
echo "Approximate round trip times in milli-seconds:"
echo "  Minimum = $min_rtt ms, Maximum = $max_rtt ms, Average = $avg_rtt ms"
