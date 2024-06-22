package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os/exec"
	"strconv"
	"strings"
)

const (
	threshold = 90
	email     = "admin@example.com"
)

func main() {
	usage, err := getDiskUsage()
	if err != nil {
		log.Fatal(err)
	}

	if usage > threshold {
		err = sendAlert(usage)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getDiskUsage() (int, error) {
	cmd := exec.Command("df", "-h", "/")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("unexpected df output")
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 5 {
		return 0, fmt.Errorf("unexpected df output format")
	}

	usageStr := strings.TrimRight(fields[4], "%")
	return strconv.Atoi(usageStr)
}

func sendAlert(usage int) error {
	from := "alert@example.com"
	password := "your-smtp-password"
	to := []string{email}
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject: Disk Usage Alert\r\n\r\nDisk usage is at %d%%, which exceeds the threshold of %d%%.", usage, threshold))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Alert sent successfully")
	return nil
}
