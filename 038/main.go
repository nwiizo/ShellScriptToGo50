package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	logDir := "/var/log/myapp"
	maxLogs := 5

	// 現在のログファイルを圧縮
	compressLog(filepath.Join(logDir, "app.log"))

	// 古いログファイルを削除
	removeOldLogs(logDir, maxLogs)

	// 新しいログファイルを作成
	createNewLog(filepath.Join(logDir, "app.log"))
}

func compressLog(logPath string) {
	input, err := os.Open(logPath)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer input.Close()

	output, err := os.Create(logPath + ".gz")
	if err != nil {
		fmt.Println("Error creating compressed file:", err)
		return
	}
	defer output.Close()

	gzipWriter := gzip.NewWriter(output)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, input)
	if err != nil {
		fmt.Println("Error compressing file:", err)
		return
	}

	os.Remove(logPath)
}

func removeOldLogs(logDir string, maxLogs int) {
	pattern := filepath.Join(logDir, "app.log.gz*")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("Error finding log files:", err)
		return
	}

	sort.Sort(sort.Reverse(sort.StringSlice(matches)))

	for i := maxLogs; i < len(matches); i++ {
		os.Remove(matches[i])
	}
}

func createNewLog(logPath string) {
	_, err := os.Create(logPath)
	if err != nil {
		fmt.Println("Error creating new log file:", err)
	}
}
