#!/bin/bash

# JWTの秘密鍵
SECRET_KEY="your_secret_key"

# ユーザー情報
sers=(
  "user1:password1"
  "user2:password2"
)

# JWTを生成する関数
generate_jwt() {
  local username=$1
  local payload="{\"username\":\"$username\"}"
  local header="{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
  local header_base64=$(echo -n "$header" | base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')
  local payload_base64=$(echo -n "$payload" | base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')
  local signature=$(echo -n "$header_base64.$payload_base64" | openssl dgst -sha256 -hmac "$SECRET_KEY" -binary | base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')
  echo "$header_base64.$payload_base64.$signature"
}

# JWTを検証する関数
verify_jwt() {
  local token=$1
  local header_payload=$(echo -n "$token" | cut -d '.' -f 1,2)
  local signature=$(echo -n "$token" | cut -d '.' -f 3)
  local expected_signature=$(echo -n "$header_payload" | openssl dgst -sha256 -hmac "$SECRET_KEY" -binary | base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')
  if [ "$signature" = "$expected_signature" ]; then
    echo "Valid"
  else
    echo "Invalid"
  fi
}

# ログインエンドポイント
login_endpoint() {
  local username=$(echo "$1" | jq -r '.username')
  local password=$(echo "$1" | jq -r '.password')
  for user in "${users[@]}"; do
    local u=$(echo "$user" | cut -d ':' -f 1)
    local p=$(echo "$user" | cut -d ':' -f 2)
    if [ "$username" = "$u" ] && [ "$password" = "$p" ]; then
      local token=$(generate_jwt "$username")
      echo "HTTP/1.1 200 OK"
      echo "Content-Type: application/json"
      echo ""
      echo "{\"token\":\"$token\"}"
      return
    fi
  done
  echo "HTTP/1.1 401 Unauthorized"
  echo "Content-Type: application/json"
  echo ""
  echo "{\"error\":\"Invalid credentials\"}"
}

# 保護されたエンドポイント
protected_endpoint() {
  local token=$1
  local validity=$(verify_jwt "$token")
  if [ "$validity" = "Valid" ]; then
    echo "HTTP/1.1 200 OK"
    echo "Content-Type: application/json"
    echo ""
    echo "{\"message\":\"Access granted\"}"
  else
    echo "HTTP/1.1 401 Unauthorized"
    echo "Content-Type: application/json"
    echo ""
    echo "{\"error\":\"Invalid token\"}"
  fi
}

# リクエストを処理する
process_request() {
  local method=$1
  local path=$2
  local headers=$3
  local data=$4
  case "$path" in
    "/login")
      if [ "$method" = "POST" ]; then
        login_endpoint "$data"
      else
        echo "HTTP/1.1 405 Method Not Allowed"
        echo "Content-Type: application/json"
        echo ""
        echo "{\"error\":\"Method not allowed\"}"
      fi
      ;;
    "/protected")
      if [ "$method" = "GET" ]; then
        local token=$(echo "$headers" | grep -i "Authorization" | cut -d ' ' -f 2)
        protected_endpoint "$token"
      else
        echo "HTTP/1.1 405 Method Not Allowed"
        echo "Content-Type: application/json"
        echo ""
        echo "{\"error\":\"Method not allowed\"}"
      fi
      ;;
    *)
      echo "HTTP/1.1 404 Not Found"
      echo "Content-Type: application/json"
      echo ""
      echo "{\"error\":\"Not found\"}"
      ;;
  esac
}

# HTTPサーバーを起動する
start_server() {
  while true; do
    local request=$(nc -l 8080)
    local method=$(echo "$request" | head -n 1 | cut -d ' ' -f 1)
    local path=$(echo "$request" | head -n 1 | cut -d ' ' -f 2)
    local headers=$(echo "$request" | sed '1,/^$/d' | sed '/^$/,$d')
    local data=$(echo "$request" | sed '1,/^$/d' | sed -n '/^$/,$p')
    process_request "$method" "$path" "$headers" "$data"
  done
}

# サーバーを起動する
start_server
