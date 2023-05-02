#!/bin/bash
# 指定した複数のURLからデータをダウンロードする。なお、並列処理を利用してください

urls=("https://sreake.com/service-sre/" "https://sreake.com/blog/5point-good-postmortem/" "https://sreake.com/blog/what-is-sre/")

download() {
  url=$1
  output=$(basename $url)
  curl -s -o $output $url
}

for url in "${urls[@]}"; do
  download $url &
done

wait

