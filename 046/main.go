package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	reportFile := "/tmp/performance_report.txt"

	report := fmt.Sprintf("Performance Report %s\n\n", time.Now().Format(time.RFC1123))

	report += "Top 5 CPU consuming processes:\n"
	report += getTopProcesses("-pcpu")

	report += "\nTop 5 Memory consuming processes:\n"
	report += getTopProcesses("-pmem")

	err := os.WriteFile(reportFile, []byte(report), 0644)
	if err != nil {
		fmt.Println("Error writing report:", err)
		return
	}

	fmt.Printf("Performance report generated at %s\n", reportFile)
}

func getTopProcesses(sortFlag string) string {
	cmd := exec.Command("ps", "aux", "--sort", sortFlag)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("Error getting process info: %v\n", err)
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 6 {
		return strings.Join(lines[:6], "\n")
	}
	return string(output)
}
