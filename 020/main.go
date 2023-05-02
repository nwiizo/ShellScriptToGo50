package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func main() {
	// config.yamlからデータベース接続情報を読み込む
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
	}

	// データベースに接続
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.User, config.Password, config.Dbname)
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// クエリ1: 全ての利用可能な書籍をタイトルと著者名で表示
	rows, _ := dbpool.Query(context.Background(), "SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id;")
	fmt.Println("全ての利用可能な書籍をタイトルと著者名で表示")
	for rows.Next() {
		var title, authorName string
		rows.Scan(&title, &authorName)
		fmt.Printf("Title: %s, Author: %s\n", title, authorName)
	}
	rows.Close()

	// クエリ2: 指定されたユーザーが借りた書籍の一覧を表示（ユーザーID: 1）
	rows, _ = dbpool.Query(context.Background(), "SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id JOIN loans ON books.id = loans.book_id WHERE loans.user_id = 1;")
	fmt.Println("\n指定されたユーザーが借りた書籍の一覧を表示（ユーザーID: 1）")
	for rows.Next() {
		var title, authorName string
		rows.Scan(&title, &authorName)
		fmt.Printf("Title: %s, Author: %s\n", title, authorName)
	}
	rows.Close()

	// クエリ3: 期限切れの貸出書籍をユーザー名とタイトルで表示
	rows, _ = dbpool.Query(context.Background(), "SELECT users.name, books.title FROM users JOIN loans ON users.id = loans.user_id JOIN books ON loans.book_id = books.id WHERE loans.due_date < CURRENT_DATE;")
	fmt.Println("\n期限切れの貸出書籍をユーザー名とタイトルで表示")
	for rows.Next() {
		var userName, title string
		rows.Scan(&userName, &title)
		fmt.Printf("User: %s, Title: %s\n", userName, title)
	}
	rows.Close()
}
