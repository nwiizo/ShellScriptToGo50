#!/bin/bash

PASSPHRASE="your-secret-passphrase"

encrypt_file() {
    openssl enc -aes-256-cbc -salt -in "$1" -out "${1}.enc" -pass pass:"$PASSPHRASE"
    rm "$1"
}

decrypt_file() {
    openssl enc -d -aes-256-cbc -in "$1" -out "${1%.enc}" -pass pass:"$PASSPHRASE"
    rm "$1"
}

encrypt_directory() {
    find "$1" -type f ! -name "*.enc" | while read file; do
        encrypt_file "$file"
    done
}

decrypt_directory() {
    find "$1" -type f -name "*.enc" | while read file; do
        decrypt_file "$file"
    done
}

case "$1" in
    encrypt)
        encrypt_directory "$2"
        ;;
    decrypt)
        decrypt_directory "$2"
        ;;
    *)
        echo "Usage: $0 {encrypt|decrypt} directory"
        exit 1
        ;;
esac
