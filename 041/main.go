package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	reportFile := "/tmp/system_report.txt"

	cpuUsage := getCPUUsage()
	memUsage := getMemoryUsage()
	diskIO := getDiskIO()

	report := fmt.Sprintf("System Report %s\n", time.Now().Format(time.RFC1123))
	report += fmt.Sprintf("CPU Usage: %.2f%%\n", cpuUsage)
	report += fmt.Sprintf("Memory Usage: %.2f%%\n", memUsage)
	report += fmt.Sprintf("Disk I/O: %.2f KB/s\n", diskIO)

	err := os.WriteFile(reportFile, []byte(report), 0644)
	if err != nil {
		fmt.Println("Error writing report:", err)
	}
}

func getCPUUsage() float64 {
	cmd := exec.Command("top", "-bn1")
	output, _ := cmd.Output()
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "Cpu(s)") {
			fields := strings.Fields(line)
			usage, _ := strconv.ParseFloat(fields[1], 64)
			return usage
		}
	}
	return 0
}

func getMemoryUsage() float64 {
	cmd := exec.Command("free")
	output, _ := cmd.Output()
	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		total, _ := strconv.ParseFloat(fields[1], 64)
		used, _ := strconv.ParseFloat(fields[2], 64)
		return (used / total) * 100
	}
	return 0
}

func getDiskIO() float64 {
	cmd := exec.Command("iostat", "-d")
	output, _ := cmd.Output()
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "sda") {
			fields := strings.Fields(line)
			io, _ := strconv.ParseFloat(fields[1], 64)
			return io
		}
	}
	return 0
}
