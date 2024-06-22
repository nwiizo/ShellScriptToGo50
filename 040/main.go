package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"
)

const (
	sourceDir = "/path/to/source"
	backupDir = "/path/to/backup"
)

func main() {
	date := time.Now().Format("20060102")

	// 最新の完全バックアップを見つける
	lastFull := findLastFullBackup()

	// 増分バックアップを作成
	backupFile := filepath.Join(backupDir, fmt.Sprintf("incr_%s.tar.gz", date))
	snapshotFile := filepath.Join(backupDir, "snapshot")

	cmd := exec.Command("tar", "-czf", backupFile, "-g", snapshotFile, "-C", filepath.Dir(sourceDir), filepath.Base(sourceDir))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error creating backup:", err)
		return
	}

	// スナップショットファイルが存在しない場合、これは完全バックアップとなる
	if _, err := os.Stat(snapshotFile); os.IsNotExist(err) {
		fullBackupFile := filepath.Join(backupDir, fmt.Sprintf("full_%s.tar.gz", date))
		err = os.Rename(backupFile, fullBackupFile)
		if err != nil {
			fmt.Println("Error renaming backup file:", err)
			return
		}
	}

	fmt.Println("Backup created successfully")
}

func findLastFullBackup() string {
	files, err := os.ReadDir(backupDir)
	if err != nil {
		fmt.Println("Error reading backup directory:", err)
		return ""
	}

	var fullBackups []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".gz" && len(file.Name()) > 5 && file.Name()[:5] == "full_" {
			fullBackups = append(fullBackups, file.Name())
		}
	}

	if len(fullBackups) == 0 {
		return ""
	}

	sort.Sort(sort.Reverse(sort.StringSlice(fullBackups)))
	return filepath.Join(backupDir, fullBackups[0])
}
