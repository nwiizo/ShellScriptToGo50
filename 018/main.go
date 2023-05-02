package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func main() {
	configFile := os.Args[1]

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return
	}

	// Note: Use "host=%s port=5432 ..." to specify the port number explicitly.
	dataSource := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", config.Host, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		fmt.Printf("Connection to PostgreSQL server at %s failed: %s\n", config.Host, err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Connection to PostgreSQL server at %s failed: %s\n", config.Host, err)
		return
	}

	fmt.Printf("Connection to PostgreSQL server at %s is successful.\n", config.Host)
}
