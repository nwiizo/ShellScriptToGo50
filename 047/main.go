package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	logFile    = "/var/log/system_update.log"
	adminEmail = "admin@example.com"
)

func main() {
	logUpdate("System update started")

	if err := updateSystem(); err != nil {
		log.Printf("Error updating system: %v", err)
	}

	if hasSecurityUpdates() {
		sendSecurityAlert()
	}

	logUpdate("System update completed")
}

func logUpdate(message string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("%s at %s", message, time.Now().Format(time.RFC1123))
}

func updateSystem() error {
	cmds := []string{
		"sudo apt update",
		"sudo apt upgrade -y",
	}

	for _, cmd := range cmds {
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("failed to execute '%s': %v", cmd, err)
		}
	}

	return nil
}

func hasSecurityUpdates() bool {
	cmd := exec.Command("sudo", "apt", "list", "--upgradable")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error checking for security updates: %v", err)
		return false
	}

	return strings.Contains(string(output), "security")
}

func sendSecurityAlert() {
	from := "alert@example.com"
	pass := "your-password"
	to := []string{adminEmail}

	msg := []byte("Subject: Security Update Alert\r\n\r\nImportant security updates are available.")

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, pass, "smtp.example.com"),
		from, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
	}
}
