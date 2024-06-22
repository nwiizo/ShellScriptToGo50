package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	logFile := "/var/log/auth.log"
	reportFile := "/tmp/failed_logins.txt"

	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	report, err := os.OpenFile(reportFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening report file:", err)
		return
	}
	defer report.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Failed password") {
			fmt.Fprintln(report, line)
			fmt.Println("Failed login attempt detected:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
