package main

import (
	"fmt"
	"net"
	"net/smtp"
	"time"
)

const (
	host     = "example.com"
	interval = 60 * time.Second
	email    = "admin@example.com"
)

func main() {
	for {
		if !checkConnection(host) {
			sendAlert(host)
		}
		time.Sleep(interval)
	}
}

func checkConnection(host string) bool {
	_, err := net.DialTimeout("tcp", host+":80", 5*time.Second)
	return err == nil
}

func sendAlert(host string) {
	from := "alert@example.com"
	pass := "your-password"
	to := []string{email}

	msg := []byte(fmt.Sprintf("Subject: Network Connection Alert\r\n\r\nConnection to %s failed at %s", host, time.Now().Format(time.RFC1123)))

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, pass, "smtp.example.com"),
		from, to, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}
