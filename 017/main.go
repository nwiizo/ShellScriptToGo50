package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	host := os.Args[1]

	if host == "" {
		fmt.Println("Usage: 017 <host>")
		return
	}
	// NOT ICMP
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "80"), time.Second*5)
	if err != nil {
		fmt.Printf("Ping to %s failed.\n", host)
		return
	}
	conn.Close()

	fmt.Printf("Ping to %s is successful.\n", host)
}
