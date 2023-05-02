package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	urls := []string{
		"https://sreake.com/service-sre/",
		"https://sreake.com/blog/5point-good-postmortem/",
		"https://sreake.com/blog/what-is-sre/",
	}

	var wg sync.WaitGroup

	// 各URLからデータを並列でダウンロードする
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			download(url)
		}(url)
	}

	wg.Wait()
}

// downloadは指定されたURLからファイルをダウンロードし、同じディレクトリに保存する
func download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	output := filepath.Base(url)
	file, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", output, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", output, err)
	}
}
