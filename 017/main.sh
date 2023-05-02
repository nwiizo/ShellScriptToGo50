#!/bin/bash

host="$1"

ping -c 1 "$host" > /dev/null 2>&1

if [ $? -eq 0 ]; then
  echo "Ping to $host is successful."
else
  echo "Ping to $host failed."
fi
