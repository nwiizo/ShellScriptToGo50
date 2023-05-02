package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
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

	// DDLを実行
	err = executeSQLFile(dbpool, "library_ddl.sql")
	if err != nil {
		log.Fatalf("Failed to execute DDL: %v", err)
	}

	// DMLを実行
	err = executeSQLFile(dbpool, "library_dml.sql")
	if err != nil {
		log.Fatalf("Failed to execute DML: %v", err)
	}

	fmt.Println("Database setup is complete.")
}

func loadConfig(filename string) (*Config, error) {
	config := &Config{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func executeSQLFile(dbpool *pgxpool.Pool, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	sql := string(data)
	_, err = dbpool.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}
