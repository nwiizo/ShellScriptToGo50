package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func processFile(filePath string) int {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return 0
	}
	// ファイルの処理ロジックをここに記述
	// 例: 単語数を数える
	wordCount := len(strings.Fields(string(content)))
	return wordCount
}

func worker(id int, jobs <-chan string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range jobs {
		wordCount := processFile(filePath)
		results <- wordCount
	}
}

func main() {
	directory := ".."
	numWorkers := 4

	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan string, len(files))
	results := make(chan int, len(files))
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, filePath := range files {
		jobs <- filePath
	}
	close(jobs)

	wg.Wait()
	close(results)

	totalWordCount := 0
	for wordCount := range results {
		totalWordCount += wordCount
	}

	fmt.Printf("Total word count: %d\n", totalWordCount)
}
