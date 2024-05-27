#!/bin/bash

directory=".."

size=$(du -sb "$directory" | awk '{print $1}')

echo "Directory size: $size bytes"
