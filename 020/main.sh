#!/bin/bash

# config.yaml からデータベース接続情報を読み込む
DB_HOST=$(grep 'host:' config.yaml | awk '{print $2}')
DB_USER=$(grep 'user:' config.yaml | awk '{print $2}')
DB_PASSWORD=$(grep 'password:' config.yaml | awk '{print $2}')
DB_NAME=$(grep 'dbname:' config.yaml | awk '{print $2}')
DB_PORT="5432"

# SQLクエリを実行
echo "全ての利用可能な書籍をタイトルと著者名で表示"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id;"

echo "指定されたユーザーが借りた書籍の一覧を表示（ユーザーID: 1）"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id JOIN loans ON books.id = loans.book_id WHERE loans.user_id = 1;"

echo "期限切れの貸出書籍をユーザー名とタイトルで表示"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT users.name, books.title FROM users JOIN loans ON users.id = loans.user_id JOIN books ON loans.book_id = books.id WHERE loans.due_date < CURRENT_DATE;"

