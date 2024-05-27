#!/bin/bash

src_dir="/path/to/source"
dst_dir="/path/to/destination"

rsync -av --delete "$src_dir/" "$dst_dir/"

echo "Directories synchronized successfully."
