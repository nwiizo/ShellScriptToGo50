package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://blog.3-shake.com/"

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Time: %s\n", duration)
}
