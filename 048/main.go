package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	dbName    = "mydb"
	backupDir = "/path/to/backups"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "restore" && len(os.Args) == 3 {
		restoreDB(os.Args[2])
	} else {
		createBackup()
	}
}

func createBackup() {
	date := time.Now().Format("20060102_150405")
	backupFile := filepath.Join(backupDir, fmt.Sprintf("%s_%s.sql", dbName, date))

	cmd := exec.Command("pg_dump", dbName)
	output, err := os.Create(backupFile)
	if err != nil {
		fmt.Printf("Error creating backup file: %v\n", err)
		return
	}
	defer output.Close()

	cmd.Stdout = output
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Backup failed: %v\n", err)
		return
	}

	fmt.Printf("Backup created successfully: %s\n", backupFile)
}

func restoreDB(backupFile string) {
	if _, err := os.Stat(backupFile); os.IsNotExist(err) {
		fmt.Printf("Backup file not found: %s\n", backupFile)
		return
	}

	cmd := exec.Command("psql", dbName)
	input, err := os.Open(backupFile)
	if err != nil {
		fmt.Printf("Error opening backup file: %v\n", err)
		return
	}
	defer input.Close()

	cmd.Stdin = input
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Restore failed: %v\n", err)
		return
	}

	fmt.Printf("Database restored from %s\n", backupFile)
}
